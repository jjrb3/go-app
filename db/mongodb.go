package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is the object connection to DB.
var MongoCN = ConnectDB()

// clientOptions is the object url configuration to MongoDB connection.
var clientOptions = options.
	Client().
	ApplyURI("mongodb+srv://JJrb3:JJrb3...@sandbox.aluia.mongodb.net/?retryWrites=true&w=majority")

// ConnectDB is the function to connect the MongoDB
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Success connection to MongoDB")

	return client
}

// CheckConnection is the function to do Ping to the MongoDB connection.
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
