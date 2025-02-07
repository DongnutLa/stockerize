package order_domain

import (
	"time"

	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Type      string             `bson:"type" json:"type"`
	Products  *[]OrderProduct    `bson:"products" json:"products"`
	User      *user_domain.User  `bson:"user" json:"user"`
	Totals    *Totals            `bson:"totals" json:"totals"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

func NewOrder(otype string, products *[]OrderProduct, user *user_domain.User, totals *Totals) *Order {
	id := primitive.NewObjectID()
	now := time.Now()

	return &Order{
		ID:        id,
		Type:      otype,
		Products:  products,
		User:      user,
		Totals:    totals,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
