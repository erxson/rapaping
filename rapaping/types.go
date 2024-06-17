package rapaping

import "time"

type ConnectionStats struct {
	Attempted int
	Connected int
	Failed    int
	MinTime   time.Duration
	MaxTime   time.Duration
	TotalTime time.Duration
}

type Info struct {
	Host string
	Port int
}

type IPInfo struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Org      string `json:"org"`
}
