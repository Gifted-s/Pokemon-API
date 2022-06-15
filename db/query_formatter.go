package db

import (
	"pokemon/m/v1/models"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetPokemonQueryFormatter takes in a query body and construct a MongoDB query and MongoDb query option from it. It returns queryFormat Struct
func GetPokemonQueryFormatter(queryBody map[string][]string) *models.GetPokemonQueryFormatterStruct {
	queryConditions := []bson.M{}
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
	}

	queryFormat := &models.GetPokemonQueryFormatterStruct{
		QueryCondition: queryConditions,
		Options:        opts,
	}
	return queryFormat
}
