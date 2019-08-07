package models

type Active struct {
	Count	int 	`json: "count"`
}

type Average struct {
	Average    float64    `json: "average"`
}

type Bandwidth struct {
	Success bool    `json:"success"`
	Units   string  `json:"units"`
	Stats   float64 `json:"stats"`
}

type NodesList struct {
	Success bool   `json:"success"`
	List    []List `json:"list"`
}
type Location struct {
	Latitude  float64 `json:"latitude"`
	City      string  `json:"city"`
	Longitude float64 `json:"longitude"`
	Country   string  `json:"country"`
}
type NetSpeed struct {
	Download float64 `json:"download"`
	Upload   float64 `json:"upload"`
}
type Load struct {
	CPU    float64 `json:"cpu"`
	Memory int     `json:"memory"`
}
type List struct {
	AccountAddr       string   `json:"account_addr"`
	IP                string   `json:"ip"`
	Latency           float64  `json:"latency"`
	VpnType           string   `json:"vpn_type"`
	Location          Location `json:"location"`
	NetSpeed          NetSpeed `json:"net_speed"`
	EncMethod         string   `json:"enc_method"`
	Moniker           string   `json:"moniker,omitempty"`
	Description       string   `json:"description,omitempty"`
	Version           string   `json:"version"`
	ActiveConnections int      `json:"active_connections"`
	Load              Load     `json:"load"`
	Rating            float64  `json:"rating,omitempty"`
	PricePerGB        float64  `json:"price_per_GB"`
}