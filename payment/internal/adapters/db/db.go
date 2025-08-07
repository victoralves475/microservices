package db

import (
	"context"
	"time"

	"github.com/victoralves475/microservices/payment/internal/application/core/domain"
	"github.com/victoralves475/microservices/payment/internal/ports"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Adapter implementa a interface ports.PaymentPort usando GORM/MySQL.
type Adapter struct {
	db *gorm.DB
}

// NewAdapter abre a conexão e retorna o Adapter.
func NewAdapter(dsn string) (*Adapter, error) {
	g, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto-migrate cria/atualiza a tabela `payments`
	if err := g.AutoMigrate(&domain.Payment{}); err != nil {
		return nil, err
	}
	return &Adapter{db: g}, nil
}

// Charge salva o pagamento e devolve o registro com o ID preenchido.
func (a *Adapter) Charge(ctx context.Context, p domain.Payment) (domain.Payment, error) {
	p.CreatedAt = time.Now().Unix()
	if err := a.db.WithContext(ctx).Create(&p).Error; err != nil {
		return domain.Payment{}, err
	}
	return p, nil
}

// Garante que Adapter implementa ports.PaymentPort
var _ ports.PaymentPort = (*Adapter)(nil)
