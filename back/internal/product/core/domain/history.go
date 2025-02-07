package product_domain

import (
	"time"

	user_domain "github.com/DongnutLa/stockio/internal/user/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type History struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Type      string             `bson:"type" json:"type"`
	Quantity  int64              `bson:"quantity" json:"quantity"`
	Price     float64            `bson:"price" json:"price"`
	Gain      float64            `bson:"gain" json:"gain"`
	Unit      string             `bson:"unit" json:"unit"`
	User      *user_domain.User  `bson:"user" json:"user"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

func NewHistory(htype string, quantity int64, price, gain float64, unit string, user *user_domain.User) *History {
	id := primitive.NewObjectID()
	now := time.Now()

	return &History{
		ID:        id,
		Type:      htype,
		Quantity:  quantity,
		Price:     price,
		Gain:      gain,
		Unit:      unit,
		User:      user,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
