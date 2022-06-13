package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateIndex () {
	pokemonCollection := GetDBCollections().PokeMons
	model := mongo.IndexModel{Keys: bson.D{primitive.E{Key: "hp", Value:  -1},  primitive.E{Key:"attack", Value:  -1}, {Key: "defense", Value:  -1} }}
	name, err := pokemonCollection.Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created: " + name)
}
