package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pokemon/m/v1/configs"
	"strconv"
	"time"
)

func GetPokeMons(queryBody map[string][]string) ([]primitive.M, error) {
	fmt.Print(queryBody)
	CreateIndex()
	opts := options.Find()
	page, _ := strconv.Atoi(queryBody["page"][0])
	limit := configs.Config.Db_Config.Fetch_Limit
	if limit != 0 {
		if page == 0 {
			page = 1
		}
		opts.SetSkip(int64((page - 1) * limit))
		opts.SetLimit(int64(limit))
	}
	queryConditions := []bson.M{}
	pokemonCollection := GetDBCollections().PokeMons
	var pokeMonInfoCursor *mongo.Cursor
	var err error
	for k, v := range queryBody {
		if k == "defense" {
			fmt.Print(k)
			queryConditions = append(queryConditions, bson.M{"defense": bson.M{queryBody[k][0]: 80}})
		}
		// if k == "attack" {
		//  queryConditions = append(queryConditions, bson.M{"attack": bson.M{queryBody["attack"][0]: 80}})
		// }
		// if k == "hp" {
		//  queryConditions = append(queryConditions, bson.M{"hp": bson.M{queryBody["hp"][0]: 80}})
		// }
		if k == "page" {
			page, _ := strconv.Atoi(queryBody[k][0])
			limit := configs.Config.Db_Config.Fetch_Limit
			if limit != 0 {
				if page == 0 {
					page = 1
				}
				opts.SetSkip(int64((page - 1) * limit))
				opts.SetLimit(int64(limit))
			}
		}
		fmt.Println("k:", k, "v:", v)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if queryBody["hp"] != nil && queryBody["defense"] != nil && queryBody["attack"] != nil {
		pokeMonInfoCursor, err = pokemonCollection.Find(ctx,
			bson.M{"$and": queryConditions},
			opts)
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
