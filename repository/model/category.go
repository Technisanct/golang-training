package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Category struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty"`
	UUID      string              `bson:"uuid"`
	Name      string              `bson:"name"`
	CreatedAt time.Time           `bson:"createdAt"`
	UpdatedAt time.Time           `bson:"updatedAt"`
}
