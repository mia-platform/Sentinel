package config

import "time"

type Configuration struct {
	ID			 string         `json:"id"`
	Output   []OutputConfig `json:"output"`
	Monitor  MonitorConfig  `json:"monitor"`
	Server   ServerConfig   `json:"server"`
	Advanced AdvancedConfig `json:"advanced"`
}

type OutputConfig struct {
	Type    string        `json:"type"`
	Webhook WebhookConfig `json:"webhook"`
	File    FileConfig    `json:"file"`
}

type WebhookConfig struct {
	URL            string         `json:"url"`
	Authentication Authentication `json:"authentication"`
}

type Authentication struct {
	Secret     SecretSource `json:"secret"`
	HeaderName string       `json:"headerName"`
}

type FileConfig struct {
	Path string `json:"path"`
}

type MonitorConfig struct {
	Interval time.Duration  `json:"interval"`
	Filters  *FiltersConfig `json:"filters"`
}
type FiltersConfig struct {
	Whitelist []string `json:"whitelist"`
	Blacklist []string `json:"blacklist"`
	Users     []string `json:"users"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

type AdvancedConfig struct {
	Debug        bool     `json:"debug"`
	OmitScanning []string `json:"omitScanning"`
}
