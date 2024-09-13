package server

import (
	"fmt"
	"follow/pkg/config"
	pb "follow/pkg/pb/connection"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(cfg config.Config, c pb.CompanyServiceServer) (*Server, error) {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return nil, err
	}

	newServer := grpc.NewServer()
	pb.RegisterCompanyServiceServer(newServer, c)

	return &Server{
		server:   newServer,
		listener: lis,
	}, nil
}

func (c *Server) Start(port string) error {
	fmt.Println("grpc server listening on port", port)
	return c.server.Serve(c.listener)
}