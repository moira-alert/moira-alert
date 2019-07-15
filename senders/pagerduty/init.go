package pagerduty

import (
	"time"

	"github.com/moira-alert/moira"
)

// Sender implements moira sender interface for pagerduty
type Sender struct {
	ImageStores  map[string]moira.ImageStore
	imageStoreID string
	logger       moira.Logger
	frontURI     string
	location     *time.Location
}

// Init loads yaml config, configures the pagerduty client
func (sender *Sender) Init(senderSettings map[string]string, logger moira.Logger, location *time.Location, dateTimeFormat string) error {
	sender.frontURI = senderSettings["front_uri"]
	sender.imageStoreID = senderSettings["image_store"]
	if sender.imageStoreID == "" {
		logger.Warningf("cannot read image_store from the config, will not be able to attach plot images to events")
	} else if !(sender.ImageStores[sender.imageStoreID].IsEnabled()) {
		logger.Warningf("image store specified (%s) has not been configured", sender.imageStoreID)
	}
	sender.logger = logger
	sender.location = location
	return nil
}
