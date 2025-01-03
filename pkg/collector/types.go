package collector

import (
	"github.com/mia-platform/sentinel/pkg/collector/environment"
	gopsutilProcesses "github.com/mia-platform/sentinel/pkg/collector/processes/gopsutil"
)

type Collector struct {
	Environment environment.EnvironmentInfo     `json:"environment"`
	Processes   []gopsutilProcesses.ProcessInfo `json:"processes"`
}
