package nginx

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/TheLazarusNetwork/LazarusTunnel/core"
	"github.com/TheLazarusNetwork/LazarusTunnel/middleware"
	"github.com/TheLazarusNetwork/LazarusTunnel/model"
	"github.com/TheLazarusNetwork/LazarusTunnel/util"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/nginx")
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
	// name := strings.ToLower(c.PostForm("name"))
	// fmt.Println("nameNginix :", name)

	reqBody := model.Tunnel{}

	c.Bind(&reqBody)

	// port allocation
	max, _ := strconv.Atoi(os.Getenv("NGINX_UPPER_RANGE"))
	min, _ := strconv.Atoi(os.Getenv("NGINX_LOWER_RANGE"))

	for {
		port, err := core.GetPort(max, min)
		fmt.Println("Port :", port)
		fmt.Println("err :", err)
		if err != nil {
			panic(err)
			// fmt.Println("this is the panic error", err)
		}

		value, msg, err := middleware.IsValidSSH(reqBody.Name, port)

		// fmt.Println("valueNginix :", value, msg)
		// // fmt.Println("caddyNginixValue :", value, msg, err)
		// fmt.Println("msg :", msg)
		// fmt.Println("err in isvalid :", err)

		if err != nil {

			fmt.Println("Error in 59 ngnx.go nginx :", err)

			resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
			c.JSON(http.StatusBadRequest, resp)
			break
		} else if value == -1 {

			if msg == "Port Already in use" {
				continue
			}

			resp = util.Message(404, msg)
			c.JSON(http.StatusBadRequest, resp)
			break
		} else if value == 1 {
			//create a tunnel struct object
			var data model.Tunnel
			data.Name = reqBody.Name
			data.Port = strconv.Itoa(port)
			data.CreatedAt = time.Now().UTC().Format(time.RFC3339)
			data.Domain = os.Getenv("NGINX_DOMAIN")

			//to add tunnel config
			err := middleware.AddSSHTunnel(data)
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
	tunnels, err := middleware.ReadSSHTunnels()
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
	tunnel, err := middleware.ReadSSHTunnel(name)
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
	tunnel, err := middleware.ReadSSHTunnel(name)
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
		err = middleware.DeleteSSHTunnel(name)
		if err != nil {
			resp = util.Message(500, "Server error, Try after some time or Contact Admin...")
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			resp = util.Message(200, "Deleted Tunnel: "+name)
			c.JSON(http.StatusOK, resp)
		}
	}

}
