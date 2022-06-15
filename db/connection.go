package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"

)

// ConnectMongoDB connects to MongoDB atlas cloud using the provided details and returns a new client
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
    return client
}
