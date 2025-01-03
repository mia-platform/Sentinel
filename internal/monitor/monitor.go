package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/mia-platform/sentinel/internal/config"
	"github.com/mia-platform/sentinel/internal/sender"
	"github.com/mia-platform/sentinel/pkg/collector"
)

func Start(ctx context.Context, cfg config.Configuration) error {
	var filters *config.FiltersConfig
	filters = nil
	ticker := time.NewTicker(cfg.Monitor.Interval)
	defer ticker.Stop()

	outputType := cfg.Output[0].Type

	fmt.Printf("Filters: %v\n", cfg.Monitor.Filters)

	if cfg.Monitor.Filters != nil {
		filters = cfg.Monitor.Filters
	}

	for {
		select {
		case <-ticker.C:
			collector, err := collector.Collect(filters)
			if err != nil {
				fmt.Printf("Error collecting data: %v\n", err)
				continue
			}

			switch outputType {
			case "webhook":
				// Invia i dati al webhook
			case "stdout":
				fmt.Println(collector)
			case "file":
				err := sender.WriteToFile(cfg.Output[0].File.Path, collector)
				if err != nil {
					fmt.Printf("Error writing to file: %v\n", err)
				}
			default:
				fmt.Printf("Output type %s not supported\n", outputType)
			}

			// Qui invierai i dati al webhook
		case <-ctx.Done():
			return nil
		}
	}
}
