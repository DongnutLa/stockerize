package order_domain

import (
	"fmt"
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

// StringToOrderType convierte un string a OrderType (con validación).
func StringToOrderType(s string) (OrderType, error) {
	switch s {
	case string(OrderTypeSale):
		return OrderTypeSale, nil
	case string(OrderTypePurchase):
		return OrderTypePurchase, nil
	default:
		return "", fmt.Errorf("invalid OrderType: %s", s)
	}
}

type PaymentMethod string

const (
	PaymentMethodCash      PaymentMethod = "CASH"
	PaymentMethodNequi     PaymentMethod = "NEQUI"
	PaymentMethodDaviplata PaymentMethod = "DAVIPLATA"
)

// StringToPaymentMethod convierte un string a PaymentMethod (con validación).
func StringToPaymentMethod(s string) (PaymentMethod, error) {
	switch s {
	case string(PaymentMethodCash):
		return PaymentMethodCash, nil
	case string(PaymentMethodNequi):
		return PaymentMethodNequi, nil
	case string(PaymentMethodDaviplata):
		return PaymentMethodDaviplata, nil
	default:
		return "", fmt.Errorf("invalid PaymentMethod: %s", s)
	}
}

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
