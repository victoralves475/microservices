package domain

import "time"

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

type Order struct {
	ID         int64       `json:"id,omitempty"`
	CustomerID int64       `json:"costumer_id"`
	OrderItems []OrderItem `json:"order_items"`
	Status     string      `json:"status"`
	CreatedAt  int64       `json:"created_at"`
}

func NewOrder(customerId int64, orderItems []OrderItem) Order {
	return Order{
		CreatedAt:  time.Now().Unix(),
		Status:     "PENDING",
		CustomerID: customerId,
		OrderItems: orderItems,
	}
}
