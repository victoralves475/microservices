package ports

import "github.com/victoalves475/microservices/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
