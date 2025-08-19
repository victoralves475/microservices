package api

import (
	"context"

	"github.com/victoralves475/microservices/payment/internal/application/core/domain"
	"github.com/victoralves475/microservices/payment/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Application implementa ports.APIPort
type Application struct {
	payment ports.PaymentPort
}

func NewApplication(payment ports.PaymentPort) *Application {
	return &Application{payment: payment}
}

func (a *Application) Charge(ctx context.Context, p domain.Payment) (domain.Payment, error) {
	if p.TotalPrice > 1000 {
		return domain.Payment{}, status.Errorf(codes.InvalidArgument, "Payment over 1000 is not allowed.")
	}
	return a.payment.Charge(ctx, p)
}
