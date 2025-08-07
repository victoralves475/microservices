package grpc

import (
	"context"
	"log"

	paymentpb "github.com/victoralves475/microservices-proto/golang/payment"
	"github.com/victoralves475/microservices/payment/internal/application/core/domain"
	"github.com/victoralves475/microservices/payment/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Adapter implementa o servidor gRPC de Payment
type Adapter struct {
	api ports.PaymentPort
	paymentpb.UnimplementedPaymentServer
}

// NewAdapter cria uma instância do Adapter com a porta de negócio fornecida.
func NewAdapter(api ports.PaymentPort) *Adapter {
	return &Adapter{api: api}
}

// Register registra o serviço Payment no servidor gRPC.
func (a *Adapter) Register(server *grpc.Server) {
	paymentpb.RegisterPaymentServer(server, a)
}

// Create recebe a requisição gRPC, executa a lógica de negócio e retorna a resposta.
func (a *Adapter) Create(ctx context.Context, req *paymentpb.CreatePaymentRequest) (*paymentpb.CreatePaymentResponse, error) {
	log.Println("Creating payment...")

	// Constrói o domínio de pagamento
	newPayment := domain.NewPayment(req.UserId, req.OrderId, req.TotalPrice)

	// Executa a cobrança na porta de negócio
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to charge: %v", err)
	}

	// Retorna o ID do pagamento criado
	return &paymentpb.CreatePaymentResponse{
		PaymentId: result.ID,
	}, nil
}
