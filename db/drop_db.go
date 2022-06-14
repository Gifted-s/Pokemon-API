package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func DropDB() (string, error) {
	pokemonCollection := GetDBCollections().PokeMons
	deleteResult, err := pokemonCollection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return "", errors.New(err.Error())
	}
	fmt.Print("Database Dropped", deleteResult)
	return "Database Dropped", nil
}
