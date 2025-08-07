package api

import (
	"context"

	"github.com/victoralves475/microservices/payment/internal/application/core/domain"
	"github.com/victoralves475/microservices/payment/internal/ports"
)

// Application é o core da lógica de Payment.
type Application struct {
	payment ports.PaymentPort
}

// NewApplication cria a instância do core com a porta injetada.
func NewApplication(payment ports.PaymentPort) *Application {
	return &Application{payment: payment}
}

// Charge é o caso de uso de criar um pagamento.
func (a *Application) Charge(ctx context.Context, p domain.Payment) (domain.Payment, error) {
	return a.payment.Charge(ctx, p)
}
