package main

import (
	"fmt"
	"os"
	"time"

	"github.com/TheLazarusNetwork/LazarusTunnel/api"
	"github.com/TheLazarusNetwork/LazarusTunnel/core"
	"github.com/TheLazarusNetwork/LazarusTunnel/middleware"
	"github.com/TheLazarusNetwork/LazarusTunnel/util"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

func init() {
	fmt.Println("Heloo")
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

	// Get Hostname for updating Log StandardFi`elds
	HostName, err := os.Hostname()
	if err != nil {
		log.Infof("Error in getting the Hostname: %v", err)
	} else {
		util.StandardFields = log.Fields{
			"hostname": HostName,
			"appname":  "TunnelAPI",
		}
	}
	// Check if loading environment variables from .env file is required
	if os.Getenv("LOAD_CONFIG_FILE") == "" {
		// Load environment variables from .env file
		err = godotenv.Load()
		if err != nil {
			log.WithFields(util.StandardFields).Fatalf("Error in reading the config file: %v", err)
		}
	}

	// initialize json files and update config files
	core.Init()
	middleware.UpdateCaddyConfig()
	middleware.UpdateNginxConfig()
}

func main() {
	log.WithFields(util.StandardFields).Infof("Starting TunnelServices Version: %s", util.Version)

	if os.Getenv("RUNTYPE") == "debug" {
		// set gin release debug
		gin.SetMode(gin.DebugMode)
	} else {
		// set gin release mode
		gin.SetMode(gin.ReleaseMode)
		// disable console color
		gin.DisableConsoleColor()
		// log level info
		log.SetLevel(log.InfoLevel)
	}
	// initialize grpc
	go api.InitGrpc()
	// creates a gin router with default middleware: logger and recovery (crash-free) middleware
	ginApp := gin.Default() //rest

	// cors middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	ginApp.Use(cors.New(config))

	// protection middleware
	ginApp.Use(helmet.Default())

	// add cache storage to gin ginApp
	ginApp.Use(func(ctx *gin.Context) {
		ctx.Set("cache", cache.New(60*time.Minute, 10*time.Minute))
		ctx.Next()
	})

	// serve static files
	ginApp.Use(static.Serve("/", static.LocalFile("./ui", false)))

	// Apply API Routes
	api.ApplyRoutes(ginApp)
	// no route redirect to frontend app
	ginApp.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": 404, "message": "Invalid Endpoint Request"})
	})

	err := ginApp.Run(fmt.Sprintf("%s:%s", os.Getenv("SERVER"), os.Getenv("PORT")))
	if err != nil {
		log.WithFields(util.StandardFields).Fatal("Failed to Start Server")
	}
}
