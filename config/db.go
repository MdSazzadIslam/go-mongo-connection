package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConnect() (*mongo.Database, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("") //enter mongodb connection string here

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil) //context.TODO when

	// it's unclear which Context to use or it is not yet available (because the

	// surrounding function has not yet been extended to accept a Context

	// parameter).

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection successfull")

	return client.Database("authentication"), nil //nil is the zero value for pointers, interfaces, maps, slices, channels and function types, representing an uninitialized value.
}
