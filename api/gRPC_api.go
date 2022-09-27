package api

import (
	"context"
	"log"
	"net"

	pb "github.com/TheLazarusNetwork/LazarusTunnel/api/v1/nginx/pb/tunnel"
	"github.com/TheLazarusNetwork/LazarusTunnel/api/v2/Caddy_gRPC"
	"github.com/TheLazarusNetwork/LazarusTunnel/api/v2/Nginix_gRPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedNginxTunnelServiceServer
}

func InitGrpc() {
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

func (s *server) GetTunnels(ctx context.Context, in *pb.Empty) (*pb.GetTunnelResponse, error) {
	tunnels, status, err := Nginix_gRPC.GetTunnels()
	return &pb.GetTunnelResponse{Message: tunnels, Status: int32(status)}, err
}

func (s *server) SetTunnel(ctx context.Context, in *pb.SetTunnelRequest) (*pb.SetTunnelResponse, error) {
	tunnel, status, err := Nginix_gRPC.SetTunnel(in.Name)
	return &pb.SetTunnelResponse{Message: tunnel, Status: int32(status)}, err
}

func (s *server) DeleteTunnel(ctx context.Context, in *pb.SetTunnelRequest) (*pb.DeleteTunnelResponse, error) {
	msg, status, err := Nginix_gRPC.DeleteTunnel(in.Name)
	return &pb.DeleteTunnelResponse{Message: msg, Status: int32(status)}, err
}

func (s *server) GetByName(ctx context.Context, in *pb.SetTunnelRequest) (*pb.SetTunnelResponse, error) {
	msg, status, err := Nginix_gRPC.GetTunnelByName(in.Name)
	return &pb.SetTunnelResponse{Message: msg, Status: int32(status)}, err
}

// caddy==================================================================

func (s *server) GetCaddyByName(ctx context.Context, in *pb.SetTunnelRequest) (*pb.SetTunnelResponse, error) {
	msg, status, err := Caddy_gRPC.CaddyGetTunnelByName(in.Name)
	return &pb.SetTunnelResponse{Message: msg, Status: int32(status)}, err
}

func (s *server) GetCaddyTunnels(ctx context.Context, in *pb.Empty) (*pb.GetTunnelResponse, error) {
	tunnels, status, err := Caddy_gRPC.CaddyGetTunnels()
	return &pb.GetTunnelResponse{Message: tunnels, Status: int32(status)}, err
}

func (s *server) SetCaddyTunnel(ctx context.Context, in *pb.SetTunnelRequest) (*pb.SetTunnelResponse, error) {
	tunnel, status, err := Caddy_gRPC.CaddySetTunnel(in.Name)
	return &pb.SetTunnelResponse{Message: tunnel, Status: int32(status)}, err
}

func (s *server) DeleteCaddyTunnel(ctx context.Context, in *pb.SetTunnelRequest) (*pb.DeleteTunnelResponse, error) {
	msg, status, err := Caddy_gRPC.CaddyDeleteTunnel(in.Name)
	return &pb.DeleteTunnelResponse{Message: msg, Status: int32(status)}, err
}
