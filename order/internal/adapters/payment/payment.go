package payment_adapter

import (
	"context"
	"github.com/victoralves475/microservices-proto/golang/payment"
	"github.com/victoralves475/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	client payment.PaymentClient
}

func NewAdapter(url string) (*Adapter, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &Adapter{client: payment.NewPaymentClient(conn)}, nil
}

func (a *Adapter) Charge(o *domain.Order) error {
	_, err := a.client.Create(context.Background(), &payment.CreatePaymentRequest{
		UserId:     o.CustomerID,
		OrderId:    o.ID,
		TotalPrice: o.TotalPrice(),
	})
	return err
}
