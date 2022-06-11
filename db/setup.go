package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupDB (queryBody map[string][]string) {
	fmt.Print(queryBody["search"][0])
	pokemonCollection := GetDBCollections().PokeMons
	model := mongo.IndexModel{Keys: bson.D{{"hp", 1},  {"attack", 1}, {"defense", 1} }}
	name, err := pokemonCollection.Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name of Index Created: " + name)
}
