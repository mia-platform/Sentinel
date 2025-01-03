package collector

import (
	"github.com/mia-platform/sentinel/internal/config"
	"github.com/mia-platform/sentinel/pkg/collector/environment"
	"github.com/mia-platform/sentinel/pkg/collector/processes"
)

func Collect(filters *config.FiltersConfig) (Collector, error) {
	environmentInfo, err := environment.GatherEnvironmentInfo()
	if err != nil {
		return Collector{}, err
	}

	processInfo, err := processes.GatherProcessInfo(filters)
	if err != nil {
		return Collector{}, err
	}

	return Collector{
		Environment: environmentInfo,
		Processes:   processInfo,
	}, nil
}
