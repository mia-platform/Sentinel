package collector

import (
	"github.com/mia-platform/sentinel/pkg/collector/environment"
	"github.com/mia-platform/sentinel/pkg/collector/processes"
)

type Collector struct {
	Environment environment.EnvironmentInfo `json:"environment"`
	Processes   []processes.ProcessInfo     `json:"processes"`
}
