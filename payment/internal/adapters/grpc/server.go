package grpc

import (
	"fmt"
	"log"
	"net"

	paymentpb "github.com/victoralves475/microservices-proto/golang/payment"
	"github.com/victoralves475/microservices/payment/config"
	"github.com/victoralves475/microservices/payment/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
	adapter    *Adapter
	port       int
}

func NewServer(api ports.APIPort, port int) *Server {
	adapter := NewAdapter(api)
	srv := grpc.NewServer()
	paymentpb.RegisterPaymentServer(srv, adapter)
	if config.GetEnv() == "development" {
		reflection.Register(srv)
	}
	return &Server{grpcServer: srv, adapter: adapter, port: port}
}

func (s *Server) Start() {
	addr := fmt.Sprintf(":%d", s.port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", addr, err)
	}
	log.Printf("Payment gRPC server listening on %s", addr)
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC on %s: %v", addr, err)
	}
}

func (s *Server) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
}
