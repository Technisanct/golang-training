package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty"`
	UUID      string              `bson:"uuid"`
	Firstname string              `bson:"firstname"`
	Lastname  string              `bson:"lastname"`
	Email     string              `bson:"email"`
	CreatedAt time.Time           `bson:"createdAt"`
}
