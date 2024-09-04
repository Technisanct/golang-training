package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	ID              *primitive.ObjectID `bson:"_id,omitempty"`
	UUID            string              `bson:"uuid"`
	Name            string              `bson:"name"`
	Price           float32             `bson:"price"`
	DiscountedPrice float32             `bson:"discountedPrice"`
	CreatedAt       time.Time           `bson:"createdAt"`
	UpdatedAt       time.Time           `bson:"updatedAt"`
}
