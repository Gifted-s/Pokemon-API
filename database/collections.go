package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Collections struct {
	PokeMons  *mongo.Collection
}

func GetDBCollections() Collections {
	dbClient := ConnectMongoDB()
	pokemons := dbClient.Database("").Collection("pokemons")
	var collections Collections
	collections.PokeMons = pokemons
	return collections
}
