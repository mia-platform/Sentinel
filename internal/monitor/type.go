package monitor

import "github.com/mia-platform/sentinel/internal/config"

type Monitor struct {
	id        string
	startTime uint64
	config    config.MonitorConfig
	output    config.OutputConfig
	status    MonitorStatus
}

type MonitorStatus string

const (
	running   MonitorStatus = "running"
	stopped   MonitorStatus = "stopped"
	exception MonitorStatus = "exception"
	idle      MonitorStatus = "idle"
)
