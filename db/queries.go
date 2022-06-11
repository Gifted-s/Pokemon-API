package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	 "time"
)

func GetPokeMons(queryBody map[string][]string) []primitive.M {
	fmt.Print(queryBody["search"][0])
	pokemonCollection := GetDBCollections().PokeMons
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	matchStage := bson.D{{"$search", bson.D{{"index", "nameSearchIndex"}, {"text",bson.M{
        "query": "Buasaur",
		"path": bson.M{
			"wildcard": "*",
		},
        "fuzzy": bson.M{},
	}}}}}
	showInfoCursor, err := pokemonCollection.Aggregate(ctx, mongo.Pipeline{matchStage})
	if err != nil {
		panic(err)
	}
	var showsWithInfo []bson.M
	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}
	return showsWithInfo
}
