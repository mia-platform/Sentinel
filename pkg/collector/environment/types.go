package environment

type EnvironmentInfo struct {
	Hostname string     `json:"hostname"`
	HostID   string     `json:"vm_id"`
	Uptime   uint64     `json:"uptime"`
	OS       string     `json:"os"`
	Timestamp string    `json:"timestamp"`
	System   SystemInfo `json:"system"`
}

type SystemInfo struct {
	CPUUsage    float64 `json:"cpuUsage"`
	MemoryUsage float64 `json:"memoryUsage"`
	DiskUsage   float64 `json:"diskUsage"`
}
