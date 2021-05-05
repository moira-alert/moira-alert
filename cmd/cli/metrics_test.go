package main

import (
	"errors"
	"testing"

	logging "github.com/moira-alert/moira/logging/zerolog_adapter"
	mocks "github.com/moira-alert/moira/mock/moira-alert"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func TestCleanupOutdatedMetrics(t *testing.T) {
	conf := getDefault()
	conf.CleanupMetrics.HotParams.CleanupBatchCount = 2
	viper.Set("hot_params", "hot_params:\n  cleanup_duration: \"-3600s\"\n  cleanup_batch: 2\n"+
		"  cleanup_batch_timeout_seconds: 1\n  cleanup_keyscan_batch: 1000")

	logger, err := logging.ConfigureLog(conf.LogFile, conf.LogLevel, "cli", conf.LogPrettyFormat)
	if err != nil {
		t.Fatal(err)
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	db := mocks.NewMockDatabase(mockCtrl)
	cursor := mocks.NewMockMetricsDatabaseCursor(mockCtrl)

	Convey("Test simple cleanup", t, func() {
		db.EXPECT().ScanMetricNames().Return(cursor)
		metricsKeys := []string{"testing.metric1"}
		cursor.EXPECT().Next().Return(metricsKeys, nil).Times(1)
		cursor.EXPECT().Next().Return(nil, errors.New("end reached")).Times(1)
		db.EXPECT().RemoveMetricsValues(gomock.Any(), gomock.Any()).Return(nil).Times(1)

		err := cleanupOutdatedMetrics(conf.CleanupMetrics, db, logger)
		So(err, ShouldBeNil)
	})

	Convey("Test batched cleanup", t, func() {
		db.EXPECT().ScanMetricNames().Return(cursor)
		metricsKeys := make([]string, 4)
		cursor.EXPECT().Next().Return(metricsKeys, nil).Times(1)
		cursor.EXPECT().Next().Return(nil, errors.New("end reached")).Times(1)
		db.EXPECT().RemoveMetricsValues(gomock.Any(), gomock.Any()).Return(nil).Times(2)

		err := cleanupOutdatedMetrics(conf.CleanupMetrics, db, logger)
		So(err, ShouldBeNil)
	})
}
