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

// type Env struct {
// 	RUNTYPE      string `env:"RUNTYPE"`
// 	SERVER       string `env:"SERVER"`
// 	PORT         int    `env:"PORT" envDefault:"9080"`
// 	APP_CONF_DIR string `env:"APP_CONF_DIR" envDefault:"./conf" envExpand:true`

// 	// # Caddy Specifications
// 	CADDY_CONF_DIR       string `env:"CADDY_CONF_DIR" envDefault:"./Caddy" envExpand:true`
// 	CADDY_INTERFACE_NAME string `env:"CADDY_INTERFACE_NAME"`
// 	CADDY_DOMAIN         string `env:"CADDY_DOMAIN"`
// 	CADDY_UPPER_RANGE    int    `env:"CADDY_UPPER_RANGE"`
// 	CADDY_LOWER_RANGE    int    `env:"CADDY_LOWER_RANGE"`

// 	// # NGINX Specifications
// 	NGINX_CONF_DIR       string `env:"CADDY_CONF_DIR" envDefault:"./NGINX" envExpand:true`
// 	NGINX_INTERFACE_NAME string `env:"NGINX_INTERFACE_NAME"`
// 	NGINX_DOMAIN         string `env:"NGINX_DOMAIN "`
// 	NGINX_UPPER_RANGE    int    `env:"NGINX_UPPER_RANGE"`
// 	NGINX_LOWER_RANGE    int    `env:"NGINX_LOWER_RANGE"`
// }
