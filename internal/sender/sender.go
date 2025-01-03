package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mia-platform/sentinel/pkg/collector"
)

func SendStatus(webhookURL string, event collector.Collector) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send status: %s", resp.Status)
	}
	return nil
}
