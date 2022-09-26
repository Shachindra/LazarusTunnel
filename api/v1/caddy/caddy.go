package caddy

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/TheLazarusNetwork/LazarusTunnel/api/v1/nginx/pb/tunnel"
	"github.com/TheLazarusNetwork/LazarusTunnel/core"
	"github.com/TheLazarusNetwork/LazarusTunnel/middleware"
	"github.com/TheLazarusNetwork/LazarusTunnel/model"
	"github.com/TheLazarusNetwork/LazarusTunnel/util"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/caddy")
	{
		g.POST("", addTunnel)
		g.GET("", getTunnels)
		g.GET(":name", getTunnel)
		g.DELETE(":name", deleteTunnel)
	}
}

var resp map[string]interface{}

//addTunnel adds new tunnel config
func addTunnel(c *gin.Context) {
	//post form parameters
	name := strings.ToLower(c.PostForm("name"))
	fmt.Println("name :", name)

	// port allocation
	max, _ := strconv.Atoi(os.Getenv("CADDY_UPPER_RANGE"))
	min, _ := strconv.Atoi(os.Getenv("CADDY_LOWER_RANGE"))

	for {
		port, err := core.GetPort(max, min)
		fmt.Println("Port caddy :", port)

		// fmt.Println("Port :", port)
		fmt.Println("err :", err)
		if err != nil {
			panic(err)

			// fmt.Println("this is the panic err :",err)

		}

		// check validity of tunnel name and port
		value, msg, err := middleware.IsValidWeb(name, port)
		fmt.Println("caddyValue :", value)
		fmt.Println("msg :", msg)
		fmt.Println("err :", err)

		// fmt.Println("value :", value)
		if err != nil {
			fmt.Println("first err != nill")
			resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
			c.JSON(http.StatusBadGateway, resp)
			break
		} else if value == -1 {
			fmt.Println("before port is already in use")
			if msg == "Port Already in use" {
				continue
			}

			resp = util.Message(404, msg)
			c.JSON(http.StatusBadRequest, resp)
			break
		} else if value == 1 {
			//create a tunnel struct object
			var data model.Tunnel
			data.Name = name
			data.Port = strconv.Itoa(port)
			data.CreatedAt = time.Now().UTC().Format(time.RFC3339)
			data.Domain = os.Getenv("CADDY_DOMAIN")

			//to add tunnel config
			err := middleware.AddWebTunnel(data)
			if err != nil {
				resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
				c.JSON(http.StatusInternalServerError, resp)
				break
			} else {
				resp = util.MessageTunnel(200, data)
				c.JSON(http.StatusOK, resp)
				break
			}
		}
	}
}

//getTunnels gets all tunnel config
func getTunnels(c *gin.Context) {
	//read all tunnel config
	tunnels, err := middleware.ReadWebTunnels()
	if err != nil {
		resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
		c.JSON(http.StatusInternalServerError, resp)
	} else {
		resp = util.MessageTunnels(200, tunnels.Tunnels)
		c.JSON(http.StatusOK, resp)
	}
}

//getTunnel get specific tunnel config
func getTunnel(c *gin.Context) {
	//get parameter
	name := c.Param("name")

	//read tunnel config
	tunnel, err := middleware.ReadWebTunnel(name)
	if err != nil {
		resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
		c.JSON(http.StatusInternalServerError, resp)
	}

	//check if tunnel exists
	if tunnel.Name == "" {
		resp = util.Message(404, "Tunnel Doesn't Exists")
		c.JSON(http.StatusNotFound, resp)
	} else {
		port, err := strconv.Atoi(tunnel.Port)
		if err != nil {
			util.LogError("string conv error: ", err)
			resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			status, err := core.ScanPort(port)
			if err != nil {
				resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
				c.JSON(http.StatusInternalServerError, resp)
			} else {
				tunnel.Status = status
				resp = util.MessageTunnel(200, *tunnel)
				c.JSON(http.StatusOK, resp)
			}
		}
	}
}

func deleteTunnel(c *gin.Context) {
	//get parameter
	name := c.Param("name")

	//read tunnel config
	tunnel, err := middleware.ReadWebTunnel(name)
	if err != nil {
		resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
		c.JSON(http.StatusInternalServerError, resp)
	}

	//check if tunnel exists
	if tunnel.Name == "" {
		resp = util.Message(400, "Tunnel Doesn't Exists")
		c.JSON(http.StatusBadRequest, resp)
	} else {
		//delete tunnel config
		err = middleware.DeleteWebTunnel(name)
		if err != nil {
			resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			resp = util.Message(200, "Deleted Tunnel: "+name)
			c.JSON(http.StatusOK, resp)
		}
	}

}

// caddy===========================================================================

func CaddyGetTunnelByName(nameIn string) (*tunnel.Tunnel, int, error) {
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

func CaddyDeleteTunnel(nameIn string) (string, int, error) {

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

func CaddyGetTunnels() ([]*tunnel.Tunnel, int, error) {
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

func CaddySetTunnel(nameIn string) (*tunnel.Tunnel, int, error) {
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
