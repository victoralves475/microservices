package ports

import (
	"context"

	"github.com/victoralves475/microservices/payment/internal/application/core/domain"
)

// PaymentPort define a porta de negócio do serviço de pagamento.
type PaymentPort interface {
	// Charge salva o pagamento e retorna o objeto com ID e timestamps preenchidos.
	Charge(ctx context.Context, p domain.Payment) (domain.Payment, error)
}
