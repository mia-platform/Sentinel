package collector

import (
	"github.com/shirou/gopsutil/disk"
)

func GetDiskUsage() (float64, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return diskStat.UsedPercent, nil
}
