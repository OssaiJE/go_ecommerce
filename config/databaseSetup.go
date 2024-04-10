package config

import (
	"context"
	"fmt"
	"log"
	"os"

	// "os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
    
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}
	if client != nil {
		return client
	}

	// Establish MongoDB connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("MongoDB Connected")
	return client
}


func DBCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("go_ecommerce_db").Collection(collectionName)
	return collection
}
