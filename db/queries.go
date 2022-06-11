package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPokeMons(queryBody map[string][]string) ([]primitive.M, error) {
	pokemonCollection := GetDBCollections().PokeMons
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Print(queryBody["search"][0])
	matchStage := bson.D{{"$search", bson.D{{"index", "nameSearchIndex"}, {"text", bson.M{
		"query": queryBody["search"][0],
		"path":  "name",
		"fuzzy": bson.M{},
	}}}}}
	pokeMonInfoCursor, err := pokemonCollection.Aggregate(ctx, mongo.Pipeline{matchStage})
	if err != nil {
		return []primitive.M{}, err
	}
	var pokemonInfo []bson.M
	if err = pokeMonInfoCursor.All(ctx, &pokemonInfo); err != nil {
		return []primitive.M{}, err
	}
	if len(pokemonInfo) == 0 {
		return []primitive.M{}, nil
	}
	return pokemonInfo, nil
}
