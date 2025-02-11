package order_domain

type CreateOrderDTO struct {
	Type          OrderType       `json:"type"`
	Products      *[]OrderProduct `json:"products"`
	Totals        *Totals         `json:"totals"`
	Discount      float64         `json:"discount"`
	PaymentMethod PaymentMethod   `json:"paymentMethod"`
}

type UpdateOrderDTO struct {
	ID            string          `json:"id"`
	Products      *[]OrderProduct `json:"products"`
	Totals        *Totals         `json:"totals"`
	Discount      float64         `json:"discount"`
	PaymentMethod PaymentMethod   `json:"paymentMethod"`
}
