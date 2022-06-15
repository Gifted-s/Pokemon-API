package db

import (
	"context"
	"fmt"

)
// DropIndexes removes all the indexes on the pokemon collection
func DropIndexes () {
	pokemonCollection := GetDBCollections().PokeMons
	_, err := pokemonCollection.Indexes().DropAll(context.TODO())
	if err != nil {
		panic(err)
	}
	fmt.Println("Indexes Dropped")
}
