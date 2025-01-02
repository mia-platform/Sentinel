package collector

import (
	"os/exec"
)

func GetProcessStatus(processName string) string {
	cmd := exec.Command("pgrep", "-x", processName)
	if err := cmd.Run(); err != nil {
		return "down"
	}
	return "running"
}
