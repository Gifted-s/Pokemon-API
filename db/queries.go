package db

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPokeMons(queryBody map[string][]string) ([]primitive.M, error) {
	var pokeMonInfoCursor *mongo.Cursor
	var err error
	pokemonCollection := GetDBCollections().PokeMons
    queryFormat:= GetPokemonQueryFormatter(queryBody)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if queryBody["hp"] != nil && queryBody["defense"] != nil && queryBody["attack"] != nil {
		pokeMonInfoCursor, err = pokemonCollection.Find(ctx,

			bson.M{"$and": queryFormat.QueryCondition},
			queryFormat.Options)
	}

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
