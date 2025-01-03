package telegraf

type ProcessInfo struct {
	Name        string  `json:"name"`
	PID         int32   `json:"pid"`
	User        string  `json:"user"`
	CPUUsage    float64 `json:"cpuUsage"`
	MemoryUsage float32 `json:"memUsage"`
	Command     string  `json:"command"`
	Status      string  `json:"status"`
}
