package db

import (
	"context"
	"pokemon/m/v1/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPokeMons(queryBody map[string][]string) ([]models.Pokemon, error) {
	pokemonCollection := GetDBCollections().PokeMons
	queryFormat := GetPokemonQueryFormatter(queryBody)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pokeMonInfoCursor, err := pokemonCollection.Find(ctx,
		bson.M{"$and": queryFormat.QueryCondition},
		queryFormat.Options)
	if err != nil {
		return []models.Pokemon{}, err
	}
	var pokemonInfo []models.Pokemon
	if err = pokeMonInfoCursor.All(ctx, &pokemonInfo); err != nil {
		return []models.Pokemon{}, err
	}
	if len(pokemonInfo) == 0 {
		return []models.Pokemon{}, nil
	}
	return pokemonInfo, nil
}

func InsertPokemons(pokemons []models.Pokemon) (*mongo.InsertManyResult, error) {
	pokemonCollection := GetDBCollections().PokeMons
	pokemonsToInterfaceSlice := []interface{}{}
	for _, p := range pokemons {
		pokemonsToInterfaceSlice = append(pokemonsToInterfaceSlice, p)
	}
	result, err := pokemonCollection.InsertMany(context.TODO(), pokemonsToInterfaceSlice)
	if err != nil {
		panic(err)
	}
	return result, nil
}
