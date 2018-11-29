package index

import (
	"github.com/blevesearch/bleve"
	"github.com/moira-alert/moira"
	"github.com/moira-alert/moira/index/mapping"
	"github.com/moira-alert/moira/metrics/graphite"
	"github.com/moira-alert/moira/metrics/graphite/go-metrics"
	"gopkg.in/tomb.v2"
)

const (
	defaultIndexBatchSize = 1000
	serviceName           = "searchIndex"
)

// Index represents Index for Bleve.Index type
type Index struct {
	index             bleve.Index
	logger            moira.Logger
	database          moira.Database
	tomb              tomb.Tomb
	metrics           *graphite.IndexMetrics
	inProgress        bool
	indexed           bool
	indexActualizedTS int64
}

// NewSearchIndex return new Index object
func NewSearchIndex(logger moira.Logger, database moira.Database) *Index {
	var err error
	newIndex := Index{
		logger:   logger,
		database: database,
	}
	newIndex.metrics = metrics.ConfigureIndexMetrics(serviceName)
	indexMapping := mapping.BuildIndexMapping(mapping.Trigger{})
	newIndex.index, err = buildIndex(indexMapping)
	if err != nil {
		return nil
	}
	return &newIndex
}

// Start initializes index. It creates new mapping and index all triggers from database
func (index *Index) Start() error {
	if index.inProgress || index.indexed {
		return nil
	}

	err := index.fillIndex()
	if err == nil {
		index.indexed = true
		index.inProgress = false
	}

	index.tomb.Go(index.runIndexActualizer)
	index.tomb.Go(index.runTriggersToReindexSweepper)
	index.tomb.Go(index.checkIndexActualizationLag)
	index.tomb.Go(index.checkIndexedTriggersCount)

	return err
}

// IsReady returns boolean value which determines if index is ready to use
func (index *Index) IsReady() bool {
	return index.indexed
}

// Stop stops checks triggers
func (index *Index) Stop() error {
	index.logger.Info("Stop search index")
	index.tomb.Kill(nil)
	return index.tomb.Wait()
}
