package payment_adapter

import (
	"context"
	"log"
	"time"

	paymentpb "github.com/victoralves475/microservices-proto/golang/payment"
	"github.com/victoralves475/microservices/order/internal/application/core/domain"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	client paymentpb.PaymentClient
}

func NewAdapter(url string) (*Adapter, error) {
	opts := []grpc.DialOption{
		// retries automáticos para erros transitórios
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
				grpc_retry.WithMax(5),
				grpc_retry.WithBackoff(grpc_retry.BackoffLinear(1*time.Second)),
			),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(url, opts...)
	if err != nil {
		return nil, err
	}
	return &Adapter{client: paymentpb.NewPaymentClient(conn)}, nil
}

func (a *Adapter) Charge(o *domain.Order) error {
	// timeout de 2s por chamada
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := a.client.Create(ctx, &paymentpb.CreatePaymentRequest{
		UserId:     o.CustomerID,
		OrderId:    o.ID,
		TotalPrice: o.TotalPrice(),
	})
	if err != nil {
		// log específico para deadline
		if status.Code(err) == codes.DeadlineExceeded {
			log.Printf("payment timeout (order_id=%d): %v", o.ID, err)
		}
		return err
	}
	return nil
}
