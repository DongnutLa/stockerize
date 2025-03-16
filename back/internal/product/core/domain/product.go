package product_domain

import (
	"time"

	store_domain "github.com/DongnutLa/stockio/internal/store/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID           primitive.ObjectID  `bson:"_id" json:"id"`
	Name         string              `bson:"name" json:"name"`
	Sku          string              `bson:"sku" json:"sku"`
	Store        *store_domain.Store `bson:"store" json:"store,omitempty"`
	Stock        *[]Stock            `bson:"stock" json:"stock,omitempty"`
	Prices       *[]Price            `bson:"prices" json:"prices"`
	Unit         ProductUnit         `bson:"unit" json:"unit"`
	CreatedAt    *time.Time          `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt    *time.Time          `bson:"updatedAt" json:"updatedAt,omitempty"`
	StockSummary *Stock              `bson:"-" json:"stockSummary"`
}

func NewProduct(
	name, sku string,
	store *store_domain.Store,
	stock *[]Stock,
	unit ProductUnit,
	prices *[]Price,
) *Product {
	id := primitive.NewObjectID()
	now := time.Now()

	return &Product{
		ID:        id,
		Name:      name,
		Sku:       sku,
		Store:     store,
		Stock:     stock,
		Unit:      unit,
		Prices:    prices,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
