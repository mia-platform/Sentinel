package interfaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendToWebhook(webhookURL string, event event) error {
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
