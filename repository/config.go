package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	DBName      string
	userMongoDB *mongo.Client
}
