package system

type SystemInfo struct {
	Hostname  string          `json:"hostname"`
	HostID    string          `json:"hostID"`
	Uptime    uint64          `json:"uptime"`
	OS        string          `json:"os"`
	Resources SystemResources `json:"system"`
}

type SystemResources struct {
	CPUUsage    float64 `json:"cpuUsage"`
	MemoryUsage float64 `json:"memoryUsage"`
	DiskUsage   float64 `json:"diskUsage"`
}
