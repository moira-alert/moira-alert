package metrics

import "fmt"

// NotifierMetrics is a collection of metrics used in notifier
type NotifierMetrics struct {
	SubsMalformed          Meter
	EventsReceived         Meter
	EventsMalformed        Meter
	EventsProcessingFailed Meter
	SendingFailed          Meter
	SendersOkMetrics       MetersCollection
	SendersFailedMetrics   MetersCollection
}

// ConfigureNotifierMetrics is notifier metrics configurator
func ConfigureNotifierMetrics(registry Registry, prefix string) *NotifierMetrics {
	return &NotifierMetrics{
		SubsMalformed:          registry.NewMeter(metricNameWithPrefix(prefix, "subs.malformed")),
		EventsReceived:         registry.NewMeter(metricNameWithPrefix(prefix, "events.received")),
		EventsMalformed:        registry.NewMeter(metricNameWithPrefix(prefix, "events.malformed")),
		EventsProcessingFailed: registry.NewMeter(metricNameWithPrefix(prefix, "events.failed")),
		SendingFailed:          registry.NewMeter(metricNameWithPrefix(prefix, "sending.failed")),
		SendersOkMetrics:       registry.NewMetersCollection(),
		SendersFailedMetrics:   registry.NewMetersCollection(),
	}
}

func metricNameWithPrefix(prefix, metric string) string {
	return fmt.Sprintf("%s.%s", prefix, metric)
}
