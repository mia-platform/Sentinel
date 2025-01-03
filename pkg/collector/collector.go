package collector

import (
	"github.com/mia-platform/sentinel/pkg/collector/environment"
	"github.com/mia-platform/sentinel/pkg/collector/processes"
)

func Collect() (Collector, error) {
	environmentInfo, err := environment.GatherEnvironmentInfo()
	if err != nil {
		return Collector{}, err
	}

	processInfo, err := processes.GatherProcessInfo()
	if err != nil {
		return Collector{}, err
	}

	return Collector{
		Environment: environmentInfo,
		Processes:   processInfo,
	}, nil
}
