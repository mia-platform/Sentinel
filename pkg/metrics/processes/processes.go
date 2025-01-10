package processes

import (
	"fmt"

	"github.com/mia-platform/sentinel/internal/config"
	"github.com/shirou/gopsutil/process"
)

// func filterProcesses(filter config.Filter) []*process.Process
func filterProcesses(filter config.FiltersConfig, procs []*process.Process) ([]*process.Process, error) {
	var filteredProcs []*process.Process

	userFilter := len(filter.Users) > 0
	whitelistFilter := len(filter.Whitelist) > 0
	blacklistFilter := len(filter.Blacklist) > 0

	// Check if both whitelist and blacklist filters are active
	if whitelistFilter && blacklistFilter {
		return nil, fmt.Errorf("both whitelist and blacklist filters are active. Only one can be active at a time.")
	}

	for _, p := range procs {
		procName, _ := p.Name()
		procUser, _ := p.Username()

		// Verifica il filtro Whitelist
		if whitelistFilter {
			whitelisted := false
			for _, name := range filter.Whitelist {
				if name == procName {
					whitelisted = true
					break
				}
			}
			if !whitelisted {
				continue // Escludi processi non nella whitelist
			}
		}

		// Verifica il filtro Blacklist
		if blacklistFilter {
			blacklisted := false
			for _, name := range filter.Blacklist {
				if name == procName {
					blacklisted = true
					break
				}
			}
			if blacklisted {
				continue // Escludi processi nella blacklist
			}
		}

		// Verifica il filtro Username
		if userFilter {
			userAllowed := false
			for _, user := range filter.Users {
				if user == procUser {
					userAllowed = true
					break
				}
			}
			if !userAllowed {
				continue // Escludi processi non dell'utente specificato
			}
		}

		// Aggiungi il processo filtrato
		filteredProcs = append(filteredProcs, p)
	}

	return filteredProcs, nil
}

func GatherProcessInfo(filter *config.FiltersConfig) ([]ProcessInfo, error) {
	procs, err := process.Processes()
	var filteredProcs []*process.Process
	if err != nil {
		return nil, err
	}

	if len(procs) == 0 {
		return nil, fmt.Errorf("no processes found")
	}

	filteredProcs = procs

	if filter != nil {
		filteredProcs, err = filterProcesses(*filter, procs)
		if err != nil {
			fmt.Printf("Error filtering processes: %v\n", err)
			return []ProcessInfo{}, err
		}
	}

	numFilteredProcs := len(filteredProcs)
	fmt.Printf("Found %d processes\n", numFilteredProcs)

	// Process filtering

	var processInfoList []ProcessInfo
	for _, p := range filteredProcs {
		name, err := p.Name()
		if err != nil {
			name = "unknown"
		}
		cpu, err := p.CPUPercent()
		if err != nil {
			cpu = 0
		}

		memory, err := p.MemoryPercent()
		if err != nil {
			memory = 0
		}
		//mem, err := p.MemoryInfo()
		status, err := p.Status()
		if err != nil {
			status = "unknown"
		}

		user, err := p.Username()
		if err != nil {
			user = "unknown"
		}
		command, err := p.Cmdline()
		if err != nil {
			command = "unknown"
		}

		processInfoList = append(processInfoList, ProcessInfo{
			PID:         p.Pid,
			Name:        name,
			CPUUsage:    cpu,
			MemoryUsage: memory,
			Status:      statusToString(status),
			User:        user,
			Command:     command,
		})
	}
	return processInfoList, nil
}

// R: Running S: Sleep T: Stop I: Idle Z: Zombie W: Wait L: Lock
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
