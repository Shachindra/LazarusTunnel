package template

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"

	"github.com/TheLazarusNetwork/LazarusTunnel/core"
	"github.com/TheLazarusNetwork/LazarusTunnel/model"
)

var (
	caddyTpl = `
# {{.Name}}, {{.Port}}, {{.CreatedAt}}
{{.Name}}.webtun.lazarus.network {
	reverse_proxy / 127.0.0.1:{{.Port}}
	log {
		output file /var/log/caddy/{{.Name}}.webtun.lazarus.network.access.log {
			roll_size 3MiB
			roll_keep 5
			roll_keep_for 48h
		}
		format console
	}
	encode gzip zstd

	tls connect@lazarus.network {
		protocols tls1.2 tls1.3
	}
}
`
	nginxTpl = `
# {{.Name}}, {{.Port}}, {{.CreatedAt}}
server {
	listen 6000;
	server_name {{.Name}}.sshtun.lazarus.network;
	
	location / {
		proxy_pass http://127.0.0.1:{{.Port}}$request_uri;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header Host $host;
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_buffering off;
		proxy_redirect off;
	}
}
`
)

//Caddy configuration file template
func CaddyConfigTempl(tunnel model.Tunnel) ([]byte, error) {
	t, err := template.New("config").Parse(caddyTpl)
	if err != nil {
		return nil, err
	}

	var tplBuff bytes.Buffer
	err = t.Execute(&tplBuff, tunnel)
	if err != nil {
		return nil, err
	}

	err = core.Writefile(filepath.Join(os.Getenv("CADDY_CONF_DIR"), os.Getenv("CADDY_INTERFACE_NAME")), tplBuff.Bytes())
	if err != nil {
		return nil, err
	}

	return tplBuff.Bytes(), nil
}

//Nginx configuration file template
func NginxConfigTempl(tunnel model.Tunnel) ([]byte, error) {
	t, err := template.New("config").Parse(nginxTpl)
	if err != nil {
		return nil, err
	}

	var tplBuff bytes.Buffer
	err = t.Execute(&tplBuff, tunnel)
	if err != nil {
		return nil, err
	}

	err = core.Writefile(filepath.Join(os.Getenv("NGINX_CONF_DIR"), os.Getenv("NGINX_INTERFACE_NAME")), tplBuff.Bytes())
	if err != nil {
		return nil, err
	}

	return tplBuff.Bytes(), nil
}
