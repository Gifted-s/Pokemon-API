package db

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson"
)

func DropDB() string {
	var err error
	pokemonCollection := GetDBCollections().PokeMons
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	deleteResult, err := pokemonCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return "Error Droppong DB" + err.Error()
	}
	fmt.Print(deleteResult)
	return "DB Dropped"
}
