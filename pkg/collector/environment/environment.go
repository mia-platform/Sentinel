package environment

import (
	"time"
)

func GatherEnvironmentInfo() (EnvironmentInfo, error) {
	hostID, err := getHostID()
	if err != nil {
		return EnvironmentInfo{}, err
	}

	hostname, err := getHostname()
	if err != nil {
		return EnvironmentInfo{}, err
	}

	uptime, err := getUptime()
	if err != nil {
		return EnvironmentInfo{}, err
	}

	osInfo, err := getOS()
	if err != nil {
		return EnvironmentInfo{}, err
	}

	systemInfo, err := gatherSystemInfo()
	if err != nil {
		return EnvironmentInfo{}, err
	}

	return EnvironmentInfo{
		HostID:    hostID,
		Hostname:  hostname,
		Uptime:    uptime,
		OS:        osInfo,
		Timestamp: time.Now().Format(time.RFC3339),
		System:    systemInfo,
	}, nil
}

func gatherSystemInfo() (SystemInfo, error) {
	cpuUsage, err := getCPUUsage()
	if err != nil {
		return SystemInfo{}, err
	}

	memoryUsage, err := getMemoryUsage()
	if err != nil {
		return SystemInfo{}, err
	}

	diskUsage, err := getDiskUsage()
	if err != nil {
		return SystemInfo{}, err
	}

	return SystemInfo{
		CPUUsage:    cpuUsage,
		MemoryUsage: memoryUsage,
		DiskUsage:   diskUsage,
	}, nil
}
