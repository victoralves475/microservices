package ports

import "github.com/victoralves475/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
	Update(o *domain.Order) error
}
