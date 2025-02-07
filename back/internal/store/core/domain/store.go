package store_domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Contact   string             `bson:"contact" json:"contact"`
	Address   string             `bson:"address" json:"address"`
	City      string             `bson:"city" json:"city"`
	Status    string             `bson:"status" json:"status"`
	CreatedAt *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updatedAt" json:"updatedAt"`
}

func NewStore(name, contact, address, city, status string) *Store {
	id := primitive.NewObjectID()
	now := time.Now()

	return &Store{
		ID:        id,
		Name:      name,
		Contact:   contact,
		Address:   address,
		City:      city,
		Status:    status,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
