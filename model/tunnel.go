package model

type Tunnel struct {
	Name      string `json:"name"`
	Port      string `json:"port"`
	CreatedAt string `json:"createdAt"`
	Domain    string `json:"domain"`
	Status    string `json:"status,omitempty"`
}

type Tunnels struct {
	Tunnels []Tunnel `json:"tunnels"`
}

// type TunnelName struct {
// 	Name string `json:"name"`
// }
