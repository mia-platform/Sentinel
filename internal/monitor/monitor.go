package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/mia-platform/sentinel/internal/config"
	"github.com/mia-platform/sentinel/internal/interfaces"
	"github.com/mia-platform/sentinel/pkg/metrics"
)

func (m *Monitor) Start(ctx context.Context) error {
	sentinelID := m.id
	filters := m.config.Filters
	ticker := time.NewTicker(m.config.Interval)
	defer ticker.Stop()

	outputType := m.output.Type

	fmt.Printf("Filters: %v\n", filters)

	for {
		select {
		case <-ticker.C:
			collector, err := metrics.Collect(filters)
			if err != nil {
				fmt.Printf("Error collecting data: %v\n", err)
				continue
			}

			switch outputType {
			case "webhook":
				event := interfaces.NewEvent(sentinelID, interfaces.SentinelMetrics, collector)
				err := interfaces.SendToWebhook(m.output.Webhook.URL, *event)
				if err != nil {
					fmt.Printf("Error sending to webhook: %v\n", err)
				}
			case "stdout":
				event := interfaces.NewEvent(sentinelID, interfaces.SentinelMetrics, collector)
				fmt.Println(event)
			case "file":
				event := interfaces.NewEvent(sentinelID, interfaces.SentinelMetrics, collector)
				err := interfaces.WriteToFile(m.output.File.Path, *event)
				if err != nil {
					fmt.Printf("Error writing to file: %v\n", err)
				}
			default:
				fmt.Printf("Output type %s not supported\n", outputType)
			}

		case <-ctx.Done():
			return nil
		}
	}
}

func New(cfg config.Configuration) Monitor {
	return Monitor{
		id:        cfg.ID,
		startTime: uint64(time.Now().Unix()),
		config:    cfg.Monitor,
		output:    cfg.Output[0],
		status:    idle,
	}
}

func (m *Monitor) Stop() error {
	m.status = stopped
	return nil
}

func (m *Monitor) Status() MonitorStatus {
	return m.status
}

func (m *Monitor) Uptime() uint64 {
	return uint64(time.Now().Unix()) - m.startTime
}

func (m *Monitor) Config() config.MonitorConfig {
	return m.config
}
