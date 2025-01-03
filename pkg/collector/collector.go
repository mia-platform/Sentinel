package collector

import (
	"github.com/mia-platform/sentinel/internal/config"
	"github.com/mia-platform/sentinel/pkg/collector/environment"
	gopsutilProcesses "github.com/mia-platform/sentinel/pkg/collector/processes/gopsutil"
	//telegrafProcesses "github.com/mia-platform/sentinel/pkg/collector/processes/telegraf"
)

func Collect(filters *config.FiltersConfig) (Collector, error) {
	environmentInfo, err := environment.GatherEnvironmentInfo()
	if err != nil {
		return Collector{}, err
	}

	processInfo, err := gopsutilProcesses.GatherProcessInfo(filters)
	if err != nil {
		return Collector{}, err
	}

	return Collector{
		Environment: environmentInfo,
		Processes:   processInfo,
	}, nil
}
