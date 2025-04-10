package order_domain

import (
	product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"
	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
)

type OrderProductDTO struct {
	ID        string                     `json:"id"`
	Name      string                     `json:"name"`
	Sku       string                     `json:"sku"`
	Quantity  float64                    `json:"quantity"`
	Price     float64                    `json:"price"`
	Cost      float64                    `json:"cost"`
	Unit      product_domain.ProductUnit `json:"unit"`
	UnitPrice product_domain.Price       `json:"unitPrice"`
}
type CreateOrderDTO struct {
	Type          OrderType          `json:"type"`
	Products      *[]OrderProductDTO `json:"products"`
	Totals        *Totals            `json:"totals"`
	Discount      float64            `json:"discount"`
	PaymentMethod PaymentMethod      `json:"paymentMethod"`
}

type UpdateOrderDTO struct {
	ID            string             `json:"id"`
	Products      *[]OrderProductDTO `json:"products"`
	Totals        *Totals            `json:"totals"`
	Discount      float64            `json:"discount"`
	PaymentMethod PaymentMethod      `json:"paymentMethod"`
}

type OrdersQueryParams struct {
	shared_domain.SearchQueryParams
	OrderType OrderType `json:"orderType"`
}
