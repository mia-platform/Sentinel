package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type VMStatus struct {
	Hostname      string  `json:"hostname"`
	VMID          string  `json:"vm_id"`
	Environment   string  `json:"environment"`
	CPUUsage      float64 `json:"cpu_usage"`
	MemoryUsage   float64 `json:"memory_usage"`
	DiskUsage     float64 `json:"disk_usage"`
	ProcessStatus string  `json:"process_status"`
	Timestamp     string  `json:"timestamp"`
}

func SendStatus(webhookURL string, status VMStatus) error {
	jsonData, err := json.Marshal(status)
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
