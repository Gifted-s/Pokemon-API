package database

import (
	"pokemon/m/v1/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

type Collections struct {
	PokeMons  *mongo.Collection
}
func GetDBCollections() Collections {
	dbClient := ConnectMongoDB()

	dbName := configs.Config.Db_Config.Db_Name
	pokeMonCollectionName := configs.Config.Db_Config.Pokemon_Collection_Name
	
	pokemons := dbClient.Database(dbName).Collection(pokeMonCollectionName)
	var collections Collections
	collections.PokeMons = pokemons
	return collections
}
