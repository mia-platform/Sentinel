package processes

import (
	"github.com/shirou/gopsutil/process"
)

func GatherProcessInfo() ([]ProcessInfo, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var processInfoList []ProcessInfo
	for _, p := range procs {
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
			Status:      status,
			User:        user,
			Command:     command,
		})
	}
	return processInfoList, nil
}
