package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/TheLazarusNetwork/LazarusTunnel/api"
	"github.com/TheLazarusNetwork/LazarusTunnel/api/v1/caddy"
	"github.com/TheLazarusNetwork/LazarusTunnel/api/v1/nginx"
	pb "github.com/TheLazarusNetwork/LazarusTunnel/api/v1/nginx/pb/tunnel"
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

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
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
	go initGrpc()
	// creates a gin router with default middleware: logger and recovery (crash-free) middleware
	ginApp := gin.Default()

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

// run gRPC server==========================================================================

type server struct {
	pb.UnimplementedNginxTunnelServiceServer
}

func initGrpc() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	ser := &server{}
	pb.RegisterNginxTunnelServiceServer(s, ser)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// gRPC services =============================================================================

func (s *server) GetTunnels(ctx context.Context, in *pb.Empty) (*pb.GetTunnelResponse, error) {
	tunnels, status, err := nginx.GetTunnels()
	return &pb.GetTunnelResponse{Message: tunnels, Status: int32(status)}, err
}

func (s *server) SetTunnel(ctx context.Context, in *pb.SetTunnelRequest) (*pb.SetTunnelResponse, error) {
	tunnel, status, err := nginx.SetTunnel(in.Name)
	return &pb.SetTunnelResponse{Message: tunnel, Status: int32(status)}, err
}

func (s *server) DeleteTunnel(ctx context.Context, in *pb.SetTunnelRequest) (*pb.DeleteTunnelResponse, error) {
	msg, status, err := nginx.DeleteTunnel(in.Name)
	return &pb.DeleteTunnelResponse{Message: msg, Status: int32(status)}, err
}

func (s *server) GetByName(ctx context.Context, in *pb.SetTunnelRequest) (*pb.SetTunnelResponse, error) {
	msg, status, err := nginx.GetTunnelByName(in.Name)
	return &pb.SetTunnelResponse{Message: msg, Status: int32(status)}, err
}

// caddy==================================================================

func (s *server) GetCaddyByName(ctx context.Context, in *pb.SetTunnelRequest) (*pb.SetTunnelResponse, error) {
	msg, status, err := caddy.CaddyGetTunnelByName(in.Name)
	return &pb.SetTunnelResponse{Message: msg, Status: int32(status)}, err
}

func (s *server) GetCaddyTunnels(ctx context.Context, in *pb.Empty) (*pb.GetTunnelResponse, error) {
	tunnels, status, err := caddy.CaddyGetTunnels()
	return &pb.GetTunnelResponse{Message: tunnels, Status: int32(status)}, err
}

func (s *server) SetCaddyTunnel(ctx context.Context, in *pb.SetTunnelRequest) (*pb.SetTunnelResponse, error) {
	tunnel, status, err := caddy.CaddySetTunnel(in.Name)
	return &pb.SetTunnelResponse{Message: tunnel, Status: int32(status)}, err
}

func (s *server) DeleteCaddyTunnel(ctx context.Context, in *pb.SetTunnelRequest) (*pb.DeleteTunnelResponse, error) {
	msg, status, err := caddy.CaddyDeleteTunnel(in.Name)
	return &pb.DeleteTunnelResponse{Message: msg, Status: int32(status)}, err
}
