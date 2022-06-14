package db

import (
	"context"
	"fmt"

)

func DropIndexes () {
	pokemonCollection := GetDBCollections().PokeMons
	_, err := pokemonCollection.Indexes().DropAll(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("Indexes Dropped")
}
