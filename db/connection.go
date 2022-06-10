package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"

)

func ConnectMongoDB() *mongo.Client {
	var dbName = os.Getenv("DB_USERNAME")
    var dbPassword = os.Getenv("DB_PASSWORD")

	dbConnectionURI :="mongodb+srv://"+dbName+":"+dbPassword+"@cluster0.jciel9m.mongodb.net/?retryWrites=true&w=majority"
	// Set client options
	clientOptions := options.Client().ApplyURI(dbConnectionURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
    return client
}
