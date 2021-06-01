package middleware

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/TheLazarusNetwork/LazarusTunnel/CaddyAPI/model"
	"github.com/TheLazarusNetwork/LazarusTunnel/CaddyAPI/template"
	"github.com/TheLazarusNetwork/LazarusTunnel/CaddyAPI/util"
)

var IsLetter = regexp.MustCompile(`^[a-z0-9]+$`).MatchString

func Init() {
	UpdateCaddyConfig()
}

// IsValid check if model is valid
func IsValid(name string, port int) (int, string, error) {
	// check if the name is empty
	if name == "" {
		return -1, "Tunnel Name is required", nil
	}

	// check the name field is between 3 to 40 chars
	if len(name) < 4 || len(name) > 12 {
		return -1, "Tunnel Name field must be between 4-12 chars", nil
	}

	tunnel, err := ReadTunnel(name)
	if err != nil {
		return -1, "", err
	} else if tunnel.Name != "" {
		return -1, "Tunnel Already exists", err
	}

	if !IsLetter(name) {
		return -1, "Tunnel Name should be Aplhanumeric", nil
	}

	return 1, "", nil
}

func ReadTunnels() (*model.Tunnels, error) {
	file, err := os.OpenFile(filepath.Join(os.Getenv("CADDY_CONF_DIR"), "server.json"), os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		util.LogError("File Open error: ", err)
		return nil, err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		util.LogError("File Read error: ", err)
		return nil, err
	}

	var tunnels model.Tunnels
	err = json.Unmarshal(b, &tunnels.Tunnels)
	if err != nil {
		util.LogError("Unmarshal json error: ", err)
		return nil, err
	}

	return &tunnels, nil
}

func ReadTunnel(tunnelName string) (*model.Tunnel, error) {
	tunnels, err := ReadTunnels()
	if err != nil {
		return nil, err
	}

	var data model.Tunnel
	for _, tunnel := range tunnels.Tunnels {
		if tunnel.Name == tunnelName {
			data.Name = tunnel.Name
			data.Port = tunnel.Port
			data.CreatedAt = tunnel.CreatedAt
			data.Status = tunnel.Status
			break
		}
	}

	return &data, nil
}

func AddTunnel(tunnel model.Tunnel) error {
	tunnels, err := ReadTunnels()
	if err != nil {
		return err
	}

	var updatedTunnels []model.Tunnel
	updatedTunnels = append(updatedTunnels, tunnels.Tunnels...)
	updatedTunnels = append(updatedTunnels, tunnel)

	inter, err := json.MarshalIndent(updatedTunnels, "", "   ")
	if err != nil {
		util.LogError("JSON Marshal error: ", err)
		return err
	}

	err = util.WriteFile(filepath.Join(os.Getenv("CADDY_CONF_DIR"), "server.json"), inter)
	if err != nil {
		util.LogError("File write error: ", err)
		return err
	}

	err = UpdateCaddyConfig()
	if err != nil {
		return err
	}

	return nil
}

func DeleteTunnel(tunnelName string) error {
	tunnels, err := ReadTunnels()
	if err != nil {
		return err
	}

	var updatedTunnels []model.Tunnel
	for _, tunnel := range tunnels.Tunnels {
		if tunnel.Name == tunnelName {
			continue
		}
		updatedTunnels = append(updatedTunnels, tunnel)
	}

	inter, err := json.MarshalIndent(updatedTunnels, "", "   ")
	if err != nil {
		util.LogError("JSON Marshal error: ", err)
		return err
	}

	err = util.WriteFile(filepath.Join(os.Getenv("CADDY_CONF_DIR"), "server.json"), inter)
	if err != nil {
		util.LogError("File write error: ", err)
		return err
	}

	err = UpdateCaddyConfig()
	if err != nil {
		return err
	}

	return nil
}

func UpdateCaddyConfig() error {
	tunnels, err := ReadTunnels()
	if err != nil {
		return err
	}

	path := filepath.Join(os.Getenv("CADDY_CONF_DIR"), os.Getenv("CADDY_INTERFACE_NAME"))
	if util.FileExists(path) {
		os.Remove(path)
	}

	for _, tunnel := range tunnels.Tunnels {
		_, err := template.ConfigTempl(tunnel)
		if err != nil {
			util.LogError("Caddyfile update error: ", err)
			return err
		}
	}

	return nil
}
