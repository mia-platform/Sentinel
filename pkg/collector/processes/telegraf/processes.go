package telegraf

import (
	"fmt"
	"github.com/influxdata/telegraf/plugins/inputs/procstat"

	"github.com/influxdata/telegraf"
)

type ProcessFilter struct {
	Users     []string
	Whitelist []string
	Blacklist []string
}

func GatherProcessInfo(filter *ProcessFilter) ([]ProcessInfo, error) {
	ps := &procstat.Procstat{
		Prefix: "procstat",
	}

	// Inizializza il plugin
	if err := ps.Init(); err != nil {
		return nil, fmt.Errorf("errore nell'inizializzazione di procstat: %v", err)
	}

	// Accumulatore per memorizzare le metriche
	acc := &SentinelAccumulator{
		Metrics: make([]telegraf.Metric, 0),
	}

	// Raccogli le metriche
	err := ps.Gather(acc); 
	if err != nil {
		return nil, fmt.Errorf("errore nella raccolta delle metriche: %v", err)
	}



	// Converti le metriche in una lista di ProcessInfo
	var processList []ProcessInfo
	for _, metric := range acc.Metrics {
		pid, _ := metric.GetField("pid")
		name, _ := metric.GetField("process_name")
		cpuUsage, _ := metric.GetField("cpu_usage")
		memoryUsage, _ := metric.GetField("memory_usage")
		status, _ := metric.GetField("status")
		user, _ := metric.GetField("user")
		cmdline, _ := metric.GetField("cmdline")

		process := ProcessInfo{
			PID:         int32(pid.(int)),
			Name:        name.(string),
			CPUUsage:    safeFloat(cpuUsage),
			MemoryUsage: safeFloat(memoryUsage),
			Status:      statusToString(status.(string)),
			User:        safeString(user),
			Command:     safeString(cmdline),
		}
		processList = append(processList, process)
	}

	// Filtra i processi se sono presenti criteri di filtro
	if filter != nil {
		processList, err = filterProcesses(filter, processList)
		if err != nil {
			return nil, err
		}
	}

	return processList, nil
}

// filterProcesses filtra i processi in base ai criteri forniti.
func filterProcesses(filter *ProcessFilter, processList []ProcessInfo) ([]ProcessInfo, error) {
	var filteredProcs []ProcessInfo

	userFilter := len(filter.Users) > 0
	whitelistFilter := len(filter.Whitelist) > 0
	blacklistFilter := len(filter.Blacklist) > 0

	// Check if both whitelist and blacklist filters are active
	if whitelistFilter && blacklistFilter {
		return nil, fmt.Errorf("both whitelist and blacklist filters are active. Only one can be active at a time.")
	}

	for _, p := range processList {
		if whitelistFilter && !contains(filter.Whitelist, p.Name) {
			continue
		}
		if blacklistFilter && contains(filter.Blacklist, p.Name) {
			continue
		}
		if userFilter && !contains(filter.Users, p.User) {
			continue
		}
		filteredProcs = append(filteredProcs, p)
	}

	return filteredProcs, nil
}

func statusToString(status string) string {
	switch status {
	case "R":
		return "Running"
	case "S":
		return "Sleep"
	case "T":
		return "Stop"
	case "I":
		return "Idle"
	case "Z":
		return "Zombie"
	case "W":
		return "Wait"
	case "L":
		return "Lock"
	default:
		return "Unknown"
	}
}
