package Nginix_gRPC

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/TheLazarusNetwork/LazarusTunnel/api/v1/nginx/pb/tunnel"
	"github.com/TheLazarusNetwork/LazarusTunnel/core"
	"github.com/TheLazarusNetwork/LazarusTunnel/middleware"
	"github.com/TheLazarusNetwork/LazarusTunnel/model"
	"github.com/TheLazarusNetwork/LazarusTunnel/util"
)

var resp map[string]interface{}

func DeleteTunnel(nameIn string) (string, int, error) {

	//read tunnel config
	msg := ""
	tunnel, err := middleware.ReadSSHTunnel(nameIn)
	if err != nil {
		return "", 500, errors.New("Server error, Try after some time or Contact Admin...")

	}

	//check if tunnel exists
	if tunnel.Name == "" {
		return "", 400, errors.New("Tunnel Doesn't Exists")
	} else {
		//delete tunnel config
		err = middleware.DeleteSSHTunnel(nameIn)
		if err != nil {
			return "", 500, errors.New("Server error, Try after some time or Contact Admin...")

		} else {
			msg = "Deleted Tunnel: " + nameIn
		}
	}
	return msg, 200, nil
}

func GetTunnelByName(nameIn string) (*tunnel.Tunnel, int, error) {
	//get parameter
	// name := c.Param("name")

	//read tunnel config
	tunnl, err := middleware.ReadSSHTunnel(nameIn)
	if err != nil {
		// resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
		// c.JSON(http.StatusInternalServerError, resp)
		return nil, 500, errors.New("Server error, Try after some time or Contact Admin...")

	}
	t := &tunnel.Tunnel{}
	//check if tunnel exists
	if tunnl.Name == "" {
		resp = util.Message(404, "Tunnel Doesn't Exists")
		// c.JSON(http.StatusNotFound, resp)
		return nil, 404, errors.New("Tunnel Doesn't Exists")

	} else {
		port, err := strconv.Atoi(tunnl.Port)
		if err != nil {
			util.LogError("string conv error: ", err)
			// resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
			// c.JSON(http.StatusInternalServerError, resp)

			return nil, 500, errors.New("Server error, Try after some time or Contact Admin...")

		} else {
			status, err := core.ScanPort(port)
			if err != nil {
				// resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
				// c.JSON(http.StatusInternalServerError, resp)
				return nil, 500, errors.New("Server error, Try after some time or Contact Admin...")

			} else {
				tunnl.Status = status
				// resp = util.MessageTunnel(200, *tunnel)
				// c.JSON(http.StatusOK, resp)
				t = &tunnel.Tunnel{
					Name:      tunnl.Name,
					Port:      tunnl.Port,
					CreatedAt: tunnl.CreatedAt,
					Domain:    tunnl.Domain,
					Status:    tunnl.Status,
				}
			}
		}
	}
	return t, 200, nil

}

func GetTunnels() ([]*tunnel.Tunnel, int, error) {
	//read all tunnel config
	tunnels, err := middleware.ReadSSHTunnels()
	if err != nil {
		err = errors.New("Server error, Try after some time or Contact Admin...")
		return nil, 500, err
	}
	tun := []*tunnel.Tunnel{}
	btunn, _ := json.Marshal(tunnels.Tunnels)
	json.Unmarshal(btunn, &tun)
	return tun, 200, nil
}

func SetTunnel(nameIn string) (*tunnel.Tunnel, int, error) {
	//post form parameters
	name := strings.ToLower(nameIn)
	t := &tunnel.Tunnel{}
	// port allocation
	max, _ := strconv.Atoi(os.Getenv("NGINX_UPPER_RANGE"))
	min, _ := strconv.Atoi(os.Getenv("NGINX_LOWER_RANGE"))

	for {
		port, err := core.GetPort(max, min)
		if err != nil {
			panic(err)
		}

		value, msg, err := middleware.IsValidSSH(name, port)
		if err != nil {
			return nil, 500, errors.New("Server error, Try after some time or Contact Admin...")
			break
		} else if value == -1 {
			if msg == "Port Already in use" {
				continue
			}
			return nil, 404, errors.New(msg)
			break
		} else if value == 1 {
			//create a tunnel struct object
			var data model.Tunnel
			data.Name = name
			data.Port = strconv.Itoa(port)
			data.CreatedAt = time.Now().UTC().Format(time.RFC3339)
			data.Domain = os.Getenv("NGINX_DOMAIN")

			//to add tunnel config
			err := middleware.AddSSHTunnel(data)
			if err != nil {
				return nil, 500, errors.New("Server error, Try after some time or Contact Admin...")
				break
			} else {
				resp = util.MessageTunnel(200, data)
				t = &tunnel.Tunnel{
					Name:      data.Name,
					Port:      data.Port,
					CreatedAt: data.CreatedAt,
					Domain:    data.Domain,
					Status:    data.Status,
				}
				break
			}
		}
	}
	return t, 200, nil
}
