package order_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderProduct struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Sku      string             `bson:"sku" json:"sku"`
	Quantity int64              `bson:"quantity" json:"quantity"`
	Unit     string             `bson:"unit" json:"unit"`
}

func NewOrderProduct(id primitive.ObjectID, name, sku, unit string, quantity int64) *OrderProduct {
	return &OrderProduct{
		ID:       id,
		Name:     name,
		Sku:      sku,
		Quantity: quantity,
		Unit:     unit,
	}
}
