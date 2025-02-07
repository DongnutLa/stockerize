package product_domain

import (
	"time"

	store_domain "github.com/DongnutLa/stockio/internal/store/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID        primitive.ObjectID  `bson:"_id" json:"id"`
	Name      string              `bson:"name" json:"name"`
	Sku       string              `bson:"sku" json:"sku"`
	Store     *store_domain.Store `bson:"store" json:"store"`
	Stock     *[]Stock            `bson:"stock" json:"stock"`
	History   *[]History          `bson:"history" json:"history"`
	CreatedAt *time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time          `bson:"updatedAt" json:"updatedAt"`
}

func NewProduct(name, sku string, store *store_domain.Store, stock *[]Stock, history *[]History) *Product {
	id := primitive.NewObjectID()
	now := time.Now()

	return &Product{
		ID:        id,
		Name:      name,
		Sku:       sku,
		Store:     store,
		Stock:     stock,
		History:   history,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
