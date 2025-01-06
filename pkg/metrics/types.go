package metrics

import (
	"github.com/mia-platform/sentinel/pkg/metrics/processes"
	"github.com/mia-platform/sentinel/pkg/metrics/system"
)

type Collector struct {
	system.SystemInfo
	Processes []processes.ProcessInfo `json:"processes"`
}
