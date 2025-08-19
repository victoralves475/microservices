package api

import (
	"github.com/victoralves475/microservices/order/internal/application/core/domain"
	"github.com/victoralves475/microservices/order/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{db: db, payment: payment}
}

func (a *Application) PlaceOrder(o domain.Order) (domain.Order, error) {
	// 1) limite total de itens (soma das quantidades)
	var totalQty int
	for _, it := range o.OrderItems {
		totalQty += int(it.Quantity)
	}
	if totalQty > 50 {
		return domain.Order{}, status.Errorf(codes.InvalidArgument, "Order has more than 50 items in total.")
	}

	// 2) status inicial
	o.Status = "Pending"
	if err := a.db.Save(&o); err != nil {
		return domain.Order{}, status.Errorf(codes.Internal, "failed to save order: %v", err)
	}

	// 3) tentativa de cobrança
	if err := a.payment.Charge(&o); err != nil {
		// se falhar (InvalidArgument do Payment ou qualquer erro), cancela
		o.Status = "Canceled"
		_ = a.db.Update(&o) // implemente Update no DBPort
		return domain.Order{}, err
	}

	// 4) sucesso
	o.Status = "Paid"
	_ = a.db.Update(&o)
	return o, nil
}
