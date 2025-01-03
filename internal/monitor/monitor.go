package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mia-platform/sentinel/pkg/collector"
	"github.com/mia-platform/sentinel/internal/config"
)

func Start(ctx context.Context, cfg config.Configuration) error {
	ticker := time.NewTicker(cfg.Monitor.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			collector, err := collector.Collect()
			if err != nil {
				fmt.Printf("Error collecting data: %v\n", err)
				continue
			}
			data, _ := json.Marshal(collector)
			fmt.Println(string(data)) // Qui invierai i dati al webhook
		case <-ctx.Done():
			return nil
		}
	}
}
