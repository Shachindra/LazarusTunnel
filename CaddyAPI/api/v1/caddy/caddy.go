package caddy

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/TheLazarusNetwork/LazarusTunnel/CaddyAPI/core"
	"github.com/TheLazarusNetwork/LazarusTunnel/CaddyAPI/middleware"
	"github.com/TheLazarusNetwork/LazarusTunnel/CaddyAPI/model"
	"github.com/TheLazarusNetwork/LazarusTunnel/CaddyAPI/util"
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

	// port allocation
	port, err := core.GetPort()
	if err != nil {
		panic(err)
	}

	value, msg, err := middleware.IsValid(name, port)
	if err != nil {
		resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
	} else if value == -1 {
		resp = util.Message(400, msg)
	} else if value == 1 {
		//create a tunnel struct object
		var data model.Tunnel
		data.Name = name
		data.Port = strconv.Itoa(port)
		data.CreatedAt = time.Now().UTC().Format(time.RFC3339)
		data.Status = "inactive"

		//to add tunnel config
		err := middleware.AddTunnel(data)
		if err != nil {
			resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
		} else {
			resp = util.MessageTunnel(200, data)
		}
	}
	c.JSON(http.StatusOK, resp)
}

//getTunnels gets all tunnel config
func getTunnels(c *gin.Context) {
	//read all tunnel config
	tunnels, err := middleware.ReadTunnels()
	if err != nil {
		resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
	} else {
		resp = util.MessageTunnels(200, tunnels.Tunnels)
	}

	c.JSON(http.StatusOK, resp)
}

//getTunnel get specific tunnel config
func getTunnel(c *gin.Context) {
	//get parameter
	name := c.Param("name")

	//read tunnel config
	tunnel, err := middleware.ReadTunnel(name)
	if err != nil {
		resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
	}

	//check if tunnel exists
	if tunnel.Name == "" {
		resp = util.Message(404, "Tunnel Doesn't Exists")
	} else {
		port, err := strconv.Atoi(tunnel.Port)
		if err != nil {
			util.LogError("string conv error: ", err)
			resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
		} else {
			status, err := core.ScanPort(port)
			if err != nil {
				resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
			} else {
				tunnel.Status = status
				resp = util.MessageTunnel(200, *tunnel)
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

func deleteTunnel(c *gin.Context) {
	//get parameter
	name := c.Param("name")

	//read tunnel config
	tunnel, err := middleware.ReadTunnel(name)
	if err != nil {
		resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
	}

	//check if tunnel exists
	if tunnel.Name == "" {
		resp = util.Message(404, "Tunnel Doesn't Exists")
	} else {
		//delete tunnel config
		err = middleware.DeleteTunnel(name)
		if err != nil {
			resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
		} else {
			resp = util.Message(200, "Deleted Tunnel: "+name)
		}
	}
	c.JSON(http.StatusOK, resp)
}
