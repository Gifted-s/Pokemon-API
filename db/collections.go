package db

import (
	"os"
	"pokemon/m/v1/configs"
	"pokemon/m/v1/models"
)

func GetDBCollections() models.Collections {
	dbClient := ConnectMongoDB()
	var dbName string
	var pokeMonCollectionName string
	environment := os.Getenv("ENV")
	
	switch environment {
	case "DEVELOPMENT":
		dbName = configs.Config.Db_Config.Db_Name
		pokeMonCollectionName = configs.Config.Db_Config.Pokemon_Collection_Name
	case "TEST":
		dbName = configs.Config.Db_Config.Db_Name_Test
		pokeMonCollectionName = configs.Config.Db_Config.Pokemon_Collection_Name_Test
	case "PRODUCTION":
		dbName = configs.Config.Db_Config.Db_Name
		pokeMonCollectionName = configs.Config.Db_Config.Pokemon_Collection_Name
	}

	pokemons := dbClient.Database(dbName).Collection(pokeMonCollectionName)
	var collections models.Collections
	collections.PokeMons = pokemons
	return collections
}
