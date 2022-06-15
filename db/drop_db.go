package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// DropDB clears all the documents that we have in the pokemons collection
func DropDB() (string, error) {
	pokemonCollection := GetDBCollections().PokeMons
	deleteResult, err := pokemonCollection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return "", errors.New(err.Error())
	}
	fmt.Print("Database Dropped", deleteResult)
	return "Database Dropped", nil
}
