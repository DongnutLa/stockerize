package product_domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductUnit string

const (
	ProductUnitPc ProductUnit = "PC" // Unidades
	ProductUnitLt ProductUnit = "LT" // Litros
	ProductUnitOz ProductUnit = "OZ" // Onzas
	ProductUnitKg ProductUnit = "KG" // Kilos
	ProductUnitLb ProductUnit = "LB" // Libras
)

type Stock struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Cost      float64            `bson:"cost" json:"cost"`
	Price     float64            `bson:"price" json:"price"`
	Quantity  float64            `bson:"quantity" json:"quantity"`
	Available float64            `bson:"available" json:"available"`
	Sold      float64            `bson:"sold" json:"sold"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt,omitempty"`
}

func NewStock(cost, price, quantity, available, sold float64) *Stock {
	id := primitive.NewObjectID()
	now := time.Now()

	return &Stock{
		ID:        id,
		Cost:      cost,
		Price:     price,
		Quantity:  quantity,
		Available: available,
		Sold:      sold,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
