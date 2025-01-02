package main

import (
	"log"
	"time"
	"os"

	"sentinel/internal/sender"
	"sentinel/pkg/collector"
)

func main() {
	webhookURL := "http://your-webhook-url"
	processName := "nginx" // Sostituisci con il nome del processo da monitorare

	hostname, _ := os.Hostname()
	vmID := "example-vm-id"     // Da personalizzare
	environment := "production" // Da personalizzare

	for {
		cpuUsage, _ := collector.GetCPUUsage()
		memoryUsage, _ := collector.GetMemoryUsage()
		diskUsage, _ := collector.GetDiskUsage()
		processStatus := collector.GetProcessStatus(processName)

		status := sender.VMStatus{
			Hostname:      hostname,
			VMID:          vmID,
			Environment:   environment,
			CPUUsage:      cpuUsage,
			MemoryUsage:   memoryUsage,
			DiskUsage:     diskUsage,
			ProcessStatus: processStatus,
			Timestamp:     time.Now().Format(time.RFC3339),
		}

		if err := sender.SendStatus(webhookURL, status); err != nil {
			log.Printf("Error sending status: %v\n", err)
		} else {
			log.Println("Status sent successfully")
		}

		time.Sleep(30 * time.Second) // Invio ogni 30 secondi
	}
}
