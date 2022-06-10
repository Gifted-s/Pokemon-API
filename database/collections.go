package database

import (
	"pokemon/m/v1/configs"
	"pokemon/m/v1/models"
)


func GetDBCollections() models.Collections {
	dbClient := ConnectMongoDB()

	dbName := configs.Config.Db_Config.Db_Name
	pokeMonCollectionName := configs.Config.Db_Config.Pokemon_Collection_Name

	pokemons := dbClient.Database(dbName).Collection(pokeMonCollectionName)
	var collections models.Collections
	collections.PokeMons = pokemons
	return collections
}
