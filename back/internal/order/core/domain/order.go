package order_domain

import (
	"time"

	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderType string

const (
	OrderTypeSale     OrderType = "SALE"
	OrderTypePurchase OrderType = "PURCHASE"
)

type PaymentMethod string

const (
	PaymentMethodCash      PaymentMethod = "CASH"
	PaymentMethodNequi     PaymentMethod = "NEQUI"
	PaymentMethodDaviplata PaymentMethod = "DAVIPLATA"
)

type Order struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	Type          OrderType          `bson:"type" json:"type"`
	Products      *[]OrderProduct    `bson:"products" json:"products"`
	User          *user_domain.User  `bson:"user" json:"user"`
	Totals        *Totals            `bson:"totals" json:"totals"`
	PaymentMethod PaymentMethod      `bson:"paymentMethod" json:"paymentMethod"`
	CreatedAt     *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt     *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

func NewOrder(
	otype OrderType,
	products *[]OrderProduct,
	user *user_domain.User,
	totals *Totals,
	paymentMethod PaymentMethod,
) *Order {
	id := primitive.NewObjectID()
	now := time.Now()

	return &Order{
		ID:            id,
		Type:          otype,
		Products:      products,
		User:          user,
		Totals:        totals,
		PaymentMethod: paymentMethod,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}
}
