package collector

import (
	"github.com/shirou/gopsutil/mem"
)

func GetMemoryUsage() (float64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.UsedPercent, nil
}
