package util

import (
	"fmt"
	"sync"
	"time"

	"code.byted.org/gopkg/metrics"
)

var (
	once          sync.Once
	metricsClient *metrics.MetricsClientV2
)

func InitMetricsV2(psm, prefix string) {
	once.Do(func() {
		lastPrefix := psm
		if prefix != "" {
			lastPrefix = fmt.Sprintf("%s.%s", psm, prefix)
		}
		metricsClient = metrics.NewDefaultMetricsClientV2(lastPrefix, true)
	})
}

func EmitStore(name string, value interface{}, tags map[string]string) {
	_ = metricsClient.EmitStore(name, value, metrics.Map2Tags(tags)...)
}

func EmitTimer(name string, duration time.Duration, tags map[string]string) {
	took := duration.Nanoseconds() / 1000 // us
	_ = metricsClient.EmitTimer(fmt.Sprintf("%s.latency", name), took, metrics.Map2Tags(tags)...)
}

func EmitTimerV2(name string, value interface{}, tags map[string]string) {
	_ = metricsClient.EmitTimer(name, value, metrics.Map2Tags(tags)...)
}

func EmitCounter(name string, value interface{}, tags map[string]string) {
	_ = metricsClient.EmitCounter(fmt.Sprintf("%s.count", name), value, metrics.Map2Tags(tags)...)
}
