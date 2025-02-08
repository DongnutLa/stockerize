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
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Cost      float64            `bson:"cost" json:"cost"`
	Price     float64            `bson:"price" json:"price"`
	Quantity  int64              `bson:"quantity" json:"quantity"`
	Available int64              `bson:"available" json:"available"`
	Sold      int64              `bson:"sold" json:"sold"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

func NewStock(
	cost, price float64,
	quantity, available, sold int64,
) *Stock {
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
