package mongodb

// import (
// 	"context"
// 	"fmt"
// 	"os"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var (
// 	MONGO_USERNAME     = "MONGO_USERNAME"
// 	MONGO_PASSWORD     = "MONGO_PASSWORD"
// 	MONGO_AUTHDATABASE = "MONGO_AUTHDATABASE"
// 	MONGO_URL          = "MONGO_URL"
// 	MONGO_USER_DB      = "MONGO_USER_DB"
// )

// func NewMongoDBConnection(
// 	ctx context.Context,
// ) (*mongo.Database, error) {
// 	mongo_uri := os.Getenv(MONGO_URL)
// 	mongo_user := os.Getenv(MONGO_USERNAME)
// 	mongo_pass := os.Getenv(MONGO_PASSWORD)
// 	mongo_auth := os.Getenv(MONGO_AUTHDATABASE)
// 	mongo_database := os.Getenv(MONGO_USER_DB)

// 	connectionString := fmt.Sprintf("mongodb://%s:%s@%s/%s",
// 		mongo_user, mongo_pass, mongo_uri, mongo_auth)

// 	client, err := mongo.Connect(
// 		ctx,
// 		options.Client().ApplyURI(connectionString))

// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := client.Ping(ctx, nil); err != nil {
// 		return nil, err
// 	}

// 	return client.Database(mongo_database), nil

// }
