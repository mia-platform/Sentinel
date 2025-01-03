package environment

import (
	"os"

	"github.com/shirou/gopsutil/host"
)

func getHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostname, nil
}

func getHostID() (string, error) {
	hostID, err := host.HostID()
	if err != nil {
		return "", err
	}
	return hostID, nil
}

func getUptime() (uint64, error) {
	uptime, err := host.Uptime()
	if err != nil {
		return 0, err
	}
	return uptime, nil
}

func getOS() (string, error) {
	platform, family, version, err := host.PlatformInformation()
	if err != nil {
		return "", err
	}
	return platform + " " + family + " " + version, nil
}