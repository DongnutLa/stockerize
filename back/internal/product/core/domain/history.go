package product_domain

import (
	"time"

	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductHistoryType string

const (
	ProductHistoryPurchase ProductHistoryType = "PURCHASE"
	ProductHistorySale     ProductHistoryType = "SALE"
	ProductHistoryIncrease ProductHistoryType = "INCREASE"
	ProductHistoryDecrease ProductHistoryType = "DECREASE"
	ProductHistoryInfo     ProductHistoryType = "INFO"
)

type History struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Type        ProductHistoryType `bson:"type" json:"type"`
	Quantity    float64            `bson:"quantity" json:"quantity"`
	Price       float64            `bson:"price" json:"price"`
	Gain        float64            `bson:"gain" json:"gain"`
	Unit        ProductUnit        `bson:"unit" json:"unit"`
	SubUnit     float64            `bson:"subUnit" json:"subUnit"`
	User        *user_domain.User  `bson:"user" json:"user"`
	ProductId   primitive.ObjectID `bson:"productId" json:"productId"`
	ProductName string             `bson:"productName" json:"productName"`
	ProductSku  string             `bson:"productSku" json:"productSku"`
	CreatedAt   *time.Time         `bson:"createdAt" json:"createdAt"`
}

func NewHistory(
	htype ProductHistoryType,
	quantity, price, gain float64,
	unit ProductUnit,
	subUnit float64,
	productId, productName, productSku string,
	user *user_domain.User,
) *History {
	prdId, _ := primitive.ObjectIDFromHex(productId)
	id := primitive.NewObjectID()
	now := time.Now()

	return &History{
		ID:          id,
		Type:        htype,
		Quantity:    quantity,
		Price:       price,
		Gain:        gain,
		Unit:        unit,
		SubUnit:     subUnit,
		User:        user,
		ProductId:   prdId,
		ProductName: productName,
		ProductSku:  productSku,
		CreatedAt:   &now,
	}
}
