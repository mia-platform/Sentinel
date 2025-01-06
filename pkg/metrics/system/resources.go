package system

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func gatherSystemResources() (SystemResources, error) {
	cpuUsage, err := getCPUUsage()
	if err != nil {
		return SystemResources{}, err
	}

	memoryUsage, err := getMemoryUsage()
	if err != nil {
		return SystemResources{}, err
	}

	diskUsage, err := getDiskUsage()
	if err != nil {
		return SystemResources{}, err
	}

	return SystemResources{
		CPUUsage:    cpuUsage,
		MemoryUsage: memoryUsage,
		DiskUsage:   diskUsage,
	}, nil
}

func getCPUUsage() (float64, error) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) > 0 {
		return percentages[0], nil
	}
	return 0, nil
}

func getDiskUsage() (float64, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return diskStat.UsedPercent, nil
}

func getMemoryUsage() (float64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.UsedPercent, nil
}