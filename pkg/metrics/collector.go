package metrics

import (
	"github.com/mia-platform/sentinel/internal/config"
	"github.com/mia-platform/sentinel/pkg/metrics/processes"
	"github.com/mia-platform/sentinel/pkg/metrics/system"
	//telegrafProcesses "github.com/mia-platform/sentinel/pkg/collector/processes/telegraf"
)

func Collect(filters *config.FiltersConfig) (Collector, error) {
	systemInfo, err := system.GatherSystemInfo()
	if err != nil {
		return Collector{}, err
	}

	processInfo, err := processes.GatherProcessInfo(filters)
	if err != nil {
		return Collector{}, err
	}

	return Collector{
		SystemInfo: systemInfo,
		Processes:  processInfo,
	}, nil
}
