package order_domain

import (
	product_domain "github.com/DongnutLa/stockio/internal/product/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderProduct struct {
	ID       primitive.ObjectID         `bson:"_id" json:"id"`
	Name     string                     `bson:"name" json:"name"`
	Sku      string                     `bson:"sku" json:"sku"`
	Quantity float64                    `bson:"quantity" json:"quantity"`
	Price    float64                    `bson:"price" json:"price"`
	Cost     float64                    `bson:"cost" json:"cost"`
	Unit     product_domain.ProductUnit `bson:"unit" json:"unit"`
}

func NewOrderProduct(
	id primitive.ObjectID,
	name, sku string,
	unit product_domain.ProductUnit,
	quantity, price, cost float64,
) *OrderProduct {
	return &OrderProduct{
		ID:       id,
		Name:     name,
		Sku:      sku,
		Quantity: quantity,
		Unit:     unit,
		Price:    price,
		Cost:     cost,
	}
}
