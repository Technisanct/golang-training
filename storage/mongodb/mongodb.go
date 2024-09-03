package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	ConnectionTimeout           = 15
	connectionStringTemplateDev = "mongodb://localhost:27017/%s"
)

// NewClient ... create mongoDB client
func NewClient(env string, username string, password string, database string, endpoint string) *mongo.Client {

	connectionURI := fmt.Sprintf(connectionStringTemplateDev, database)

	//if env == "prod" {
	//	connectionURI = fmt.Sprintf(connectionStringTemplate, username, password, endpoint, database)
	//} else if env == "stg" {
	//	connectionURI = fmt.Sprintf(connectionStringTemplate, username, password, endpoint, database)
	//}

	mongoDBClient, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		panic(fmt.Sprintf("failed to create mongoDB Client. Err: %v", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeout*time.Second)
	defer cancel()

	if err = mongoDBClient.Connect(ctx); err != nil {
		panic(fmt.Sprintf("failed to connect to mongoDB. Err: %v", err))
	}

	//Force a connection to verify our connection string
	if err = mongoDBClient.Ping(ctx, nil); err != nil {
		panic(fmt.Sprintf("Failed to ping cluster: %v", err))
	}

	return mongoDBClient
}
