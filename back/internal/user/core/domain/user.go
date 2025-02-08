package user_domain

import (
	"time"

	store_domain "github.com/DongnutLa/stockio/internal/store/core/domain"
	"github.com/DongnutLa/stockio/internal/zshared/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStatus string

const (
	UserActive   UserStatus = "ACTIVE"
	UserInactive UserStatus = "INACTIVE"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Status    UserStatus         `bson:"status" json:"status"`
	Role      constants.UserRole `bson:"role" json:"role"`
	Store     store_domain.Store `bson:"store" json:"store"`
	Timestamp *time.Time         `bson:"timestamp" json:"timestamp"`
}

func NewUser(
	name, email string,
	status UserStatus,
	role constants.UserRole,
	store store_domain.Store,
) *User {
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

type CreateUserDTO struct {
	Name   string             `json:"name"`
	Email  string             `json:"email"`
	Store  store_domain.Store `json:"store"`
	Status string             `json:"status"`
}
