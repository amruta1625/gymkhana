package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/amruta1625/gymkhana/controllers"
)

func main() {
	// Replace <password> with your actual password
	password := "amruta"
	uri := "mongodb+srv://swecha:" + password + "@cluster0.4o1cvll.mongodb.net/your_database_name_here?retryWrites=true&w=majority"

	// Establish MongoDB client
	client := getClient(uri)
	defer client.Disconnect(context.Background())

	// Access the specific database you want to work with
	database := client.Database("swecha")

	r := httprouter.New()
	uc := controllers.NewUserController(database)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getClient(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	return client
}
