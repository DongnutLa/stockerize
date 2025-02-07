package user_domain

import (
	"time"

	store_domain "github.com/DongnutLa/stockio/internal/store/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Status    string             `bson:"status" json:"status"`
	Role      string             `bson:"role" json:"role"`
	Store     store_domain.Store `bson:"store" json:"store"`
	Timestamp *time.Time         `bson:"timestamp" json:"timestamp"`
}

func NewUser(name, email, status, role string, store store_domain.Store) *User {
	id := primitive.NewObjectID()
	now := time.Now()

	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Status:    status,
		Role:      role,
		Store:     store,
		Timestamp: &now,
	}
}
