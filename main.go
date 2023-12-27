package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Replace <password> with the actual password for the swecha623 user and URL encode it
	password := "Mummy#1625"
	encodedPassword := url.QueryEscape(password)

	// Construct the connection string with the encoded password
	connectionURI := fmt.Sprintf("mongodb+srv://swecha623:%s@cluster0.4o1cvll.mongodb.net/",
		encodedPassword)

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Close the connection
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
