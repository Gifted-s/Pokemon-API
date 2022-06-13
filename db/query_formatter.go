package db

import (
	"pokemon/m/v1/configs"
	"pokemon/m/v1/models"
	"strconv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPokemonQueryFormatter(queryBody map[string][]string) *models.GetPokemonQueryFormatterStruct {
	queryConditions := []bson.M{}
	var page int
	var limit int
	opts := options.Find()
	for k, v := range queryBody {
		if k == "defense" {
			valToInteger,_ := strconv.Atoi(v[1])
			queryConditions = append(queryConditions, bson.M{"defense": bson.M{v[0]: valToInteger}})
		}
		if k == "attack" {
			valToInteger,_ := strconv.Atoi(v[1])
			queryConditions = append(queryConditions, bson.M{"attack": bson.M{v[0]: valToInteger}})
		}
		if k == "hp" {
			valToInteger,_ := strconv.Atoi(v[1])
			queryConditions = append(queryConditions, bson.M{"hp": bson.M{v[0]: valToInteger}})
		}
		if k == "page" {
			page, _ = strconv.Atoi(v[0])
			limit = configs.Config.Db_Config.Fetch_Limit
			if limit != 0 {
				if page == 0 {
					page = 1
				}
				opts.SetSkip(int64((page - 1) * limit))
				opts.SetLimit(int64(limit))
			}
		}
	}

	queryFormat := &models.GetPokemonQueryFormatterStruct{
		QueryCondition: queryConditions,
		Options:        opts,
	}
	return queryFormat
}
