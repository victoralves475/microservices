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

// Server encapsula o servidor gRPC de Payment.
type Server struct {
	grpcServer *grpc.Server
	adapter    *Adapter
	port       int
}

// NewServer monta o servidor gRPC, registra o handler e, se em dev, habilita reflection.
func NewServer(paymentPort ports.PaymentPort, port int) *Server {
	adapter := NewAdapter(paymentPort)

	srv := grpc.NewServer()
	paymentpb.RegisterPaymentServer(srv, adapter)

	if config.GetEnv() == "development" {
		reflection.Register(srv)
	}

	return &Server{
		grpcServer: srv,
		adapter:    adapter,
		port:       port,
	}
}

// Start inicia o servidor na porta configurada.
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

// Stop encerra graciosamente o servidor gRPC.
func (s *Server) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
}
