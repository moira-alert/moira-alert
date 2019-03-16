package redis

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/op/go-logging"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/moira-alert/moira"
	"github.com/moira-alert/moira/database"
)

func TestLastCheck(t *testing.T) {
	logger, _ := logging.GetLogger("dataBase")
	dataBase := newTestDatabase(logger, config)
	dataBase.flush()
	defer dataBase.flush()
	var triggerMaintenanceTS int64

	Convey("LastCheck manipulation", t, func() {
		Convey("Test read write delete", func() {
			triggerID := uuid.Must(uuid.NewV4()).String()
			err := dataBase.SetTriggerLastCheck(triggerID, &lastCheckTest, false)
			So(err, ShouldBeNil)

			actual, err := dataBase.GetTriggerLastCheck(triggerID)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, lastCheckTest)

			err = dataBase.RemoveTriggerLastCheck(triggerID)
			So(err, ShouldBeNil)

			actual, err = dataBase.GetTriggerLastCheck(triggerID)
			So(err, ShouldResemble, database.ErrNil)
			So(actual, ShouldResemble, moira.CheckData{})
		})

		Convey("Test no lastcheck", func() {
			triggerID := uuid.Must(uuid.NewV4()).String()
			actual, err := dataBase.GetTriggerLastCheck(triggerID)
			So(err, ShouldBeError)
			So(err, ShouldResemble, database.ErrNil)
			So(actual, ShouldResemble, moira.CheckData{})
		})

		Convey("Test set metrics check maintenance", func() {
			Convey("While no check", func() {
				triggerID := uuid.Must(uuid.NewV4()).String()
				err := dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{}, nil, "",0)
				So(err, ShouldBeNil)
			})

			Convey("While no metrics", func() {
				triggerID := uuid.Must(uuid.NewV4()).String()
				err := dataBase.SetTriggerLastCheck(triggerID, &lastCheckWithNoMetrics, false)
				So(err, ShouldBeNil)

				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{"metric1": 1, "metric5": 5}, nil, "", 0)
				So(err, ShouldBeNil)

				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, lastCheckWithNoMetrics)
			})

			Convey("While no metrics to change", func() {
				triggerID := uuid.Must(uuid.NewV4()).String()
				err := dataBase.SetTriggerLastCheck(triggerID, &lastCheckTest, false)
				So(err, ShouldBeNil)

				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{"metric11": 1, "metric55": 5}, nil, "",0)
				So(err, ShouldBeNil)

				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, lastCheckTest)
			})

			Convey("Has metrics to change", func() {
				checkData := lastCheckTest
				triggerID := uuid.Must(uuid.NewV4()).String()
				userLogin := "test"
				var timeCallMaintenance = int64(3)


				err := dataBase.SetTriggerLastCheck(triggerID, &checkData, false)
				So(err, ShouldBeNil)

				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{"metric1": 1, "metric5": 5}, nil, userLogin, timeCallMaintenance)
				So(err, ShouldBeNil)
				metric1 := checkData.Metrics["metric1"]
				metric5 := checkData.Metrics["metric5"]

				metric1.MaintenanceWho = moira.MaintenanceWho{}
				metric5.MaintenanceWho = moira.MaintenanceWho{}
				metric1.Maintenance = 1
				metric5.Maintenance = 5
				metric1.MaintenanceWho.StopMaintenanceUser = &userLogin
				metric1.MaintenanceWho.StopMaintenanceTime = &timeCallMaintenance
				metric5.MaintenanceWho.StartMaintenanceUser = &userLogin
				metric5.MaintenanceWho.StartMaintenanceTime = &timeCallMaintenance

				checkData.Metrics["metric1"] = metric1
				checkData.Metrics["metric5"] = metric5

				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, checkData)
			})
		})

		Convey("Test set Trigger and metrics check maintenance", func() {
			Convey("While no check", func() {
				triggerID := uuid.Must(uuid.NewV4()).String()
				err := dataBase.SetTriggerCheckMaintenance(triggerID, make(map[string]int64), nil, "", 0)
				So(err, ShouldBeNil)
			})

			Convey("Set metrics maintenance while no metrics", func() {
				triggerID := uuid.Must(uuid.NewV4()).String()
				err := dataBase.SetTriggerLastCheck(triggerID, &lastCheckWithNoMetrics, false)
				So(err, ShouldBeNil)

				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{"metric1": 1, "metric5": 5}, nil,"", 0)
				So(err, ShouldBeNil)

				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, lastCheckWithNoMetrics)
			})

			Convey("Set trigger maintenance while no metrics", func() {
				triggerID := uuid.Must(uuid.NewV4()).String()
				err := dataBase.SetTriggerLastCheck(triggerID, &lastCheckWithNoMetrics, false)
				So(err, ShouldBeNil)

				triggerMaintenanceTS = 1000

				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{"metric1": 1, "metric5": 5}, &triggerMaintenanceTS, "",0)
				So(err, ShouldBeNil)

				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, lastCheckWithNoMetricsWithMaintenance)
			})

			Convey("Set metrics maintenance while no metrics to change", func() {
				triggerID := uuid.Must(uuid.NewV4()).String()
				err := dataBase.SetTriggerLastCheck(triggerID, &lastCheckTest, false)
				So(err, ShouldBeNil)

				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{"metric11": 1, "metric55": 5}, nil,"",0)
				So(err, ShouldBeNil)

				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, lastCheckTest)
			})

			Convey("Set trigger maintenance while no metrics to change", func() {
				newLastCheckTest := lastCheckTest
				newLastCheckTest.Maintenance = 1000
				triggerID := uuid.Must(uuid.NewV4()).String()
				err := dataBase.SetTriggerLastCheck(triggerID, &lastCheckTest, false)
				So(err, ShouldBeNil)

				triggerMaintenanceTS = 1000
				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{"metric11": 1, "metric55": 5}, &triggerMaintenanceTS,"",0 )
				So(err, ShouldBeNil)

				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, newLastCheckTest)
			})

			Convey("Set metrics maintenance while has metrics to change", func() {
				checkData := lastCheckTest
				triggerID := uuid.Must(uuid.NewV4()).String()
				userLogin := "test"
				var timeCallMaintenance = int64(3)

				err := dataBase.SetTriggerLastCheck(triggerID, &checkData, false)
				So(err, ShouldBeNil)

				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{"metric1": 1, "metric5": 5}, nil, userLogin, timeCallMaintenance)
				So(err, ShouldBeNil)
				metric1 := checkData.Metrics["metric1"]
				metric5 := checkData.Metrics["metric5"]

				metric1.MaintenanceWho = moira.MaintenanceWho{}
				metric5.MaintenanceWho = moira.MaintenanceWho{}
				metric1.Maintenance = 1
				metric5.Maintenance = 5
				metric1.MaintenanceWho.StopMaintenanceUser = &userLogin
				metric1.MaintenanceWho.StopMaintenanceTime = &timeCallMaintenance
				metric5.MaintenanceWho.StartMaintenanceUser = &userLogin
				metric5.MaintenanceWho.StartMaintenanceTime = &timeCallMaintenance

				checkData.Metrics["metric1"] = metric1
				checkData.Metrics["metric5"] = metric5

				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, checkData)
			})

			Convey("Set trigger and metrics maintenance while has metrics to change", func() {
				checkData := lastCheckTest
				triggerID := uuid.Must(uuid.NewV4()).String()
				checkData.MaintenanceWho = moira.MaintenanceWho{}
				userLogin := "test"
				var timeCallMaintenance = int64(3)
				err := dataBase.SetTriggerLastCheck(triggerID, &checkData, false)
				So(err, ShouldBeNil)

				triggerMaintenanceTS = 1000
				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{"metric1": 1, "metric5": 5}, &triggerMaintenanceTS, userLogin, timeCallMaintenance)
				So(err, ShouldBeNil)
				metric1 := checkData.Metrics["metric1"]
				metric5 := checkData.Metrics["metric5"]

				metric1.MaintenanceWho = moira.MaintenanceWho{}
				metric5.MaintenanceWho = moira.MaintenanceWho{}
				metric1.Maintenance = 1
				metric5.Maintenance = 5
				metric1.MaintenanceWho.StopMaintenanceUser = &userLogin
				metric1.MaintenanceWho.StopMaintenanceTime = &timeCallMaintenance
				metric5.MaintenanceWho.StartMaintenanceUser = &userLogin
				metric5.MaintenanceWho.StartMaintenanceTime = &timeCallMaintenance

				checkData.Metrics["metric1"] = metric1
				checkData.Metrics["metric5"] = metric5
				checkData.Maintenance = 1000
				checkData.MaintenanceWho.StartMaintenanceUser = &userLogin
				checkData.MaintenanceWho.StartMaintenanceTime = &timeCallMaintenance


				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, checkData)
			})

			Convey("Set trigger maintenance to 0 and metrics maintenance", func() {
				checkData := lastCheckTest
				triggerID := uuid.Must(uuid.NewV4()).String()
				err := dataBase.SetTriggerLastCheck(triggerID, &checkData, false)
				So(err, ShouldBeNil)

				triggerMaintenanceTS = 0
				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{}, &triggerMaintenanceTS,"",0)
				So(err, ShouldBeNil)
				checkData.Maintenance = 0

				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, checkData)
			})
		})

		Convey("Test last check manipulations update 'triggers to reindex' list", func() {
			dataBase.flush()
			triggerID := uuid.Must(uuid.NewV4()).String()

			// there was no trigger with such ID, so function should return true
			So(dataBase.checkDataScoreChanged(triggerID, &lastCheckWithNoMetrics), ShouldBeTrue)

			// set new last check. Should add a trigger to a reindex set
			err := dataBase.SetTriggerLastCheck(triggerID, &lastCheckWithNoMetrics, false)
			So(err, ShouldBeNil)

			So(dataBase.checkDataScoreChanged(triggerID, &lastCheckWithNoMetrics), ShouldBeFalse)

			So(dataBase.checkDataScoreChanged(triggerID, &lastCheckTest), ShouldBeTrue)

			actual, err := dataBase.FetchTriggersToReindex(time.Now().Unix() - 1)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []string{triggerID})

			time.Sleep(time.Second)

			err = dataBase.SetTriggerLastCheck(triggerID, &lastCheckTest, false)
			So(err, ShouldBeNil)

			actual, err = dataBase.FetchTriggersToReindex(time.Now().Unix() - 10)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []string{triggerID})

			err = dataBase.RemoveTriggersToReindex(time.Now().Unix() + 10)
			So(err, ShouldBeNil)

			actual, err = dataBase.FetchTriggersToReindex(time.Now().Unix() - 10)
			So(err, ShouldBeNil)
			So(actual, ShouldBeEmpty)

			err = dataBase.RemoveTriggerLastCheck(triggerID)
			So(err, ShouldBeNil)

			actual, err = dataBase.FetchTriggersToReindex(time.Now().Unix() - 1)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, []string{triggerID})
		})
	})
}

func TestRemoteLastCheck(t *testing.T) {
	logger, _ := logging.GetLogger("dataBase")
	dataBase := newTestDatabase(logger, config)
	dataBase.flush()
	defer dataBase.flush()

	Convey("LastCheck manipulation", t, func() {
		Convey("Test read write delete", func() {
			triggerID := uuid.Must(uuid.NewV4()).String()
			err := dataBase.SetTriggerLastCheck(triggerID, &lastCheckTest, true)
			So(err, ShouldBeNil)

			actual, err := dataBase.GetTriggerLastCheck(triggerID)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, lastCheckTest)

			err = dataBase.RemoveTriggerLastCheck(triggerID)
			So(err, ShouldBeNil)

			actual, err = dataBase.GetTriggerLastCheck(triggerID)
			So(err, ShouldResemble, database.ErrNil)
			So(actual, ShouldResemble, moira.CheckData{})
		})

		Convey("Test no lastcheck", func() {
			triggerID := uuid.Must(uuid.NewV4()).String()
			actual, err := dataBase.GetTriggerLastCheck(triggerID)
			So(err, ShouldBeError)
			So(err, ShouldResemble, database.ErrNil)
			So(actual, ShouldResemble, moira.CheckData{})
		})

		Convey("Test set trigger check maintenance", func() {
			Convey("While no check", func() {
				triggerID := uuid.Must(uuid.NewV4()).String()
				err := dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{}, nil,"",0)
				So(err, ShouldBeNil)
			})

			Convey("While no metrics", func() {
				triggerID := uuid.Must(uuid.NewV4()).String()
				err := dataBase.SetTriggerLastCheck(triggerID, &lastCheckWithNoMetrics, true)
				So(err, ShouldBeNil)

				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{"metric1": 1, "metric5": 5}, nil,"",0)
				So(err, ShouldBeNil)

				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, lastCheckWithNoMetrics)
			})

			Convey("While no metrics to change", func() {
				triggerID := uuid.Must(uuid.NewV4()).String()
				err := dataBase.SetTriggerLastCheck(triggerID, &lastCheckTest, true)
				So(err, ShouldBeNil)

				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{"metric11": 1, "metric55": 5}, nil,"",0)
				So(err, ShouldBeNil)

				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, lastCheckTest)
			})

			Convey("Has metrics to change", func() {
				checkData := lastCheckTest
				triggerID := uuid.Must(uuid.NewV4()).String()
				userLogin := "test"
				var timeCallMaintenance = int64(3)

				err := dataBase.SetTriggerLastCheck(triggerID, &checkData, true)
				So(err, ShouldBeNil)

				err = dataBase.SetTriggerCheckMaintenance(triggerID, map[string]int64{"metric1": 1, "metric5": 5}, nil, userLogin, timeCallMaintenance)
				So(err, ShouldBeNil)
				metric1 := checkData.Metrics["metric1"]
				metric5 := checkData.Metrics["metric5"]

				metric1.MaintenanceWho = moira.MaintenanceWho{}
				metric5.MaintenanceWho = moira.MaintenanceWho{}
				metric1.Maintenance = 1
				metric5.Maintenance = 5
				metric1.MaintenanceWho.StopMaintenanceUser = &userLogin
				metric1.MaintenanceWho.StopMaintenanceTime = &timeCallMaintenance
				metric5.MaintenanceWho.StartMaintenanceUser = &userLogin
				metric5.MaintenanceWho.StartMaintenanceTime = &timeCallMaintenance

				checkData.Metrics["metric1"] = metric1
				checkData.Metrics["metric5"] = metric5

				actual, err := dataBase.GetTriggerLastCheck(triggerID)
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, checkData)
			})
		})
	})
}

func TestLastCheckErrorConnection(t *testing.T) {
	logger, _ := logging.GetLogger("dataBase")
	dataBase := newTestDatabase(logger, emptyConfig)
	dataBase.flush()
	defer dataBase.flush()
	Convey("Should throw error when no connection", t, func() {
		actual1, err := dataBase.GetTriggerLastCheck("123")
		So(actual1, ShouldResemble, moira.CheckData{})
		So(err, ShouldNotBeNil)

		err = dataBase.SetTriggerLastCheck("123", &lastCheckTest, false)
		So(err, ShouldNotBeNil)

		err = dataBase.RemoveTriggerLastCheck("123")
		So(err, ShouldNotBeNil)

		var triggerMaintenanceTS int64 = 123
		err = dataBase.SetTriggerCheckMaintenance("123", map[string]int64{}, &triggerMaintenanceTS,"",0)
		So(err, ShouldNotBeNil)

		actual2, err := dataBase.GetTriggerLastCheck("123")
		So(actual2, ShouldResemble, moira.CheckData{})
		So(err, ShouldNotBeNil)
	})
}
func TestSetMaintenanceUserAndTime (t *testing.T) {
	startMaintenanceUser := "testStartMtUser"
	startMaintenanceTime := int64(1550304140)
	stopMaintenanceUser := "testStopMtUser"
	stopMaintenanceTime := int64(1553068940)
	triggerMaintenanceTS := int64(1552723340)

	Convey("Test trigger", t, func() {
		Convey("User anonymous", func() {
			actual := lastCheckTest
			setMaintenanceUserAndTime(&lastCheckTest, &triggerMaintenanceTS, "anonymous", startMaintenanceTime)
			actual.MaintenanceWho.StartMaintenanceUser = nil
			actual.MaintenanceWho.StartMaintenanceTime = nil
			actual.MaintenanceWho.StopMaintenanceUser = nil
			actual.MaintenanceWho.StopMaintenanceTime = nil
			So(actual, ShouldResemble, lastCheckTest)
		})
		Convey("User '' ", func() {
			actual := lastCheckTest
			setMaintenanceUserAndTime(&lastCheckTest, &triggerMaintenanceTS, "", startMaintenanceTime)
			actual.MaintenanceWho.StartMaintenanceUser = nil
			actual.MaintenanceWho.StartMaintenanceTime = nil
			actual.MaintenanceWho.StopMaintenanceUser = nil
			actual.MaintenanceWho.StopMaintenanceTime = nil
			So(actual, ShouldResemble, lastCheckTest)
		})
		Convey("User and time start maintenance", func() {
			actual := lastCheckTest
			setMaintenanceUserAndTime(&lastCheckTest, &triggerMaintenanceTS, startMaintenanceUser, startMaintenanceTime)
			actual.MaintenanceWho.StartMaintenanceUser = &startMaintenanceUser
			actual.MaintenanceWho.StartMaintenanceTime = &startMaintenanceTime
			actual.MaintenanceWho.StopMaintenanceUser = nil
			actual.MaintenanceWho.StopMaintenanceTime = nil
			So(actual, ShouldResemble, lastCheckTest)
		})
		Convey("User and time stop maintenance", func() {
			actual := lastCheckTest
			setMaintenanceUserAndTime(&lastCheckTest, &triggerMaintenanceTS, stopMaintenanceUser, stopMaintenanceTime)
			actual.MaintenanceWho.StartMaintenanceUser = nil
			actual.MaintenanceWho.StartMaintenanceTime = nil
			actual.MaintenanceWho.StopMaintenanceUser = &stopMaintenanceUser
			actual.MaintenanceWho.StopMaintenanceTime = &stopMaintenanceTime
			So(actual, ShouldResemble, lastCheckTest)
		})
		Convey("User and time start maintenance if set user and time stop maintenance", func() {
			actual := lastCheckTest
			lastCheckTest.MaintenanceWho.StopMaintenanceUser = &stopMaintenanceUser
			lastCheckTest.MaintenanceWho.StopMaintenanceTime = &stopMaintenanceTime
			setMaintenanceUserAndTime(&lastCheckTest, &triggerMaintenanceTS, startMaintenanceUser, startMaintenanceTime)
			actual.MaintenanceWho.StartMaintenanceUser = &startMaintenanceUser
			actual.MaintenanceWho.StartMaintenanceTime = &startMaintenanceTime
			actual.MaintenanceWho.StopMaintenanceUser = nil
			actual.MaintenanceWho.StopMaintenanceTime = nil
			So(actual, ShouldResemble, lastCheckTest)
		})
	})

	Convey("Test metric", t, func() {
		Convey("User anonymous", func(){
			actual := lastMetricsTest
			setMaintenanceUserAndTime(&lastMetricsTest, &triggerMaintenanceTS, "anonymous", startMaintenanceTime)
			actual.MaintenanceWho.StartMaintenanceUser = nil
			actual.MaintenanceWho.StartMaintenanceTime = nil
			actual.MaintenanceWho.StopMaintenanceUser = nil
			actual.MaintenanceWho.StopMaintenanceTime = nil
			So(actual, ShouldResemble, lastMetricsTest)
		})
		Convey("User '' ", func() {
			actual := lastMetricsTest
			setMaintenanceUserAndTime(&lastMetricsTest, &triggerMaintenanceTS, "", startMaintenanceTime)
			actual.MaintenanceWho.StartMaintenanceUser = nil
			actual.MaintenanceWho.StartMaintenanceTime = nil
			actual.MaintenanceWho.StopMaintenanceUser = nil
			actual.MaintenanceWho.StopMaintenanceTime = nil
			So(actual, ShouldResemble, lastMetricsTest)
		})
		Convey("User and time start maintenance", func(){
			actual := lastMetricsTest
			setMaintenanceUserAndTime(&lastMetricsTest, &triggerMaintenanceTS, startMaintenanceUser, startMaintenanceTime)
			actual.MaintenanceWho.StartMaintenanceUser = &startMaintenanceUser
			actual.MaintenanceWho.StartMaintenanceTime = &startMaintenanceTime
			actual.MaintenanceWho.StopMaintenanceUser = nil
			actual.MaintenanceWho.StopMaintenanceTime = nil
			So(actual, ShouldResemble, lastMetricsTest)
		})
		Convey("User and time stop maintenance", func(){
			actual := lastMetricsTest
			setMaintenanceUserAndTime(&lastMetricsTest, &triggerMaintenanceTS, stopMaintenanceUser, stopMaintenanceTime)
			actual.MaintenanceWho.StartMaintenanceUser = nil
			actual.MaintenanceWho.StartMaintenanceTime = nil
			actual.MaintenanceWho.StopMaintenanceUser = &stopMaintenanceUser
			actual.MaintenanceWho.StopMaintenanceTime = &stopMaintenanceTime
			So(actual, ShouldResemble, lastMetricsTest)
		})
		Convey("User and time start maintenance if set user and time stop maintenance", func(){
			actual := lastMetricsTest
			lastCheckTest.MaintenanceWho.StopMaintenanceUser = &stopMaintenanceUser
			lastCheckTest.MaintenanceWho.StopMaintenanceTime = &stopMaintenanceTime
			setMaintenanceUserAndTime(&lastMetricsTest, &triggerMaintenanceTS, startMaintenanceUser, startMaintenanceTime)
			actual.MaintenanceWho.StartMaintenanceUser = &startMaintenanceUser
			actual.MaintenanceWho.StartMaintenanceTime = &startMaintenanceTime
			actual.MaintenanceWho.StopMaintenanceUser = nil
			actual.MaintenanceWho.StopMaintenanceTime = nil
			So(actual, ShouldResemble, lastMetricsTest)
		})
	})
}


var lastCheckTest = moira.CheckData{
	Score:     6000,
	State:     moira.StateOK,
	Timestamp: 1504509981,
	Metrics: map[string]moira.MetricState{
		"metric1": {
			EventTimestamp: 1504449789,
			State:          moira.StateNODATA,
			Suppressed:     false,
			Timestamp:      1504509380,
		},
		"metric2": {
			EventTimestamp: 1504449789,
			State:          moira.StateNODATA,
			Suppressed:     false,
			Timestamp:      1504509380,
		},
		"metric3": {
			EventTimestamp: 1504449789,
			State:          moira.StateNODATA,
			Suppressed:     false,
			Timestamp:      1504509380,
		},
		"metric4": {
			EventTimestamp: 1504463770,
			State:          moira.StateNODATA,
			Suppressed:     false,
			Timestamp:      1504509380,
		},
		"metric5": {
			EventTimestamp: 1504463770,
			State:          moira.StateNODATA,
			Suppressed:     false,
			Timestamp:      1504509380,
		},
		"metric6": {
			EventTimestamp: 1504463770,
			State:          "Ok",
			Suppressed:     false,
			Timestamp:      1504509380,
		},
	},
}

var lastMetricsTest = moira.MetricState {
	EventTimestamp: 1504449789,
	State:          moira.StateNODATA,
	Suppressed:     false,
	Timestamp:      1504509380,
}

var lastCheckWithNoMetrics = moira.CheckData{
	Score:     0,
	State:     moira.StateOK,
	Timestamp: 1504509981,
	Metrics:   make(map[string]moira.MetricState),
}

var lastCheckWithNoMetricsWithMaintenance = moira.CheckData{
	Score:       0,
	State:       moira.StateOK,
	Timestamp:   1504509981,
	Maintenance: 1000,
	Metrics:     make(map[string]moira.MetricState),
}
