package system

func GatherSystemInfo() (SystemInfo, error) {
	hostID, err := getHostID()
	if err != nil {
		return SystemInfo{}, err
	}

	hostname, err := getHostname()
	if err != nil {
		return SystemInfo{}, err
	}

	systemUptime, err := getUptime()
	if err != nil {
		return SystemInfo{}, err
	}

	osInfo, err := getOS()
	if err != nil {
		return SystemInfo{}, err
	}

	systemResources, err := gatherSystemResources()
	if err != nil {
		return SystemInfo{}, err
	}

	return SystemInfo{
		HostID:    hostID,
		Hostname:  hostname,
		Uptime:    systemUptime,
		OS:        osInfo,
		Resources: systemResources,
	}, nil
}
