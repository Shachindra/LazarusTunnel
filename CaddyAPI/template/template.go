package template

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"

	"github.com/TheLazarusNetwork/LazarusTunnel/CaddyAPI/core"
	"github.com/TheLazarusNetwork/LazarusTunnel/CaddyAPI/model"
)

var (
	configTpl = `
# {{.Name}}, {{.Port}}, {{.CreatedAt}}
{{.Name}}.lazarus.network {
	reverse_proxy / 127.0.0.1:{{.Port}}
	log {
		output file /var/log/caddy/{{.Name}}.lazarus.network.access.log {
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
)

// Templ
func ConfigTempl(tunnel model.Tunnel) ([]byte, error) {
	t, err := template.New("config").Parse(configTpl)
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
