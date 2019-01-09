package notifier

import (
	"time"

	"github.com/moira-alert/moira/remote"
)

// Config is sending settings including log settings
type Config struct {
	Enabled          bool
	SendingTimeout   time.Duration
	ResendingTimeout time.Duration
	Senders          []map[string]string
	LogFile          string
	LogLevel         string
	FrontURL         string
	Location         *time.Location
	DateTimeFormat   string
	RemoteConfig     *remote.Config
}
