package api

import (
	"github.com/victoralves475/microservices/order/internal/application/core/domain"
	"github.com/victoralves475/microservices/order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{db: db, payment: payment}
}

func (a Application) PlaceOrder(o domain.Order) (domain.Order, error) {
	if err := a.db.Save(&o); err != nil {
		return domain.Order{}, err
	}
	if err := a.payment.Charge(&o); err != nil {
		return domain.Order{}, err
	}
	return o, nil
}
