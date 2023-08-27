package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnection() {

	username := "admin"
	password := "admin"
	host := "localhost"
	port := 27017
	authDatabase := "admin"
	// databaseName := "users"

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		username, password, host, port, authDatabase)

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	err = client.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexão encerrada.")

}
