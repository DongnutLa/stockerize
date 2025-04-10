package order_domain

import (
	"time"

	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	"github.com/samber/lo"
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
	Consecutive   string             `bson:"consecutive" json:"consecutive"`
	Products      *[]OrderProduct    `bson:"products" json:"products"`
	User          *user_domain.User  `bson:"user" json:"user"`
	Totals        *Totals            `bson:"totals" json:"totals"`
	PaymentMethod PaymentMethod      `bson:"paymentMethod" json:"paymentMethod"`
	CreatedAt     *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt     *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

func NewOrder(
	otype OrderType,
	productsDto *[]OrderProductDTO,
	user *user_domain.User,
	totals *Totals,
	paymentMethod PaymentMethod,
	consecutive string,
) *Order {
	id := primitive.NewObjectID()
	now := time.Now()

	products := lo.Map(*productsDto, dto2OrderProduct)
	return &Order{
		ID:            id,
		Type:          otype,
		Consecutive:   consecutive,
		Products:      &products,
		User:          user,
		Totals:        totals,
		PaymentMethod: paymentMethod,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}
}

func dto2OrderProduct(product OrderProductDTO, _ int) OrderProduct {
	id, _ := primitive.ObjectIDFromHex(product.ID)
	return *NewOrderProduct(
		id,
		product.Name,
		product.Sku,
		product.Unit,
		product.Quantity,
		product.Price,
		product.Cost,
	)
}
