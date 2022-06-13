package helpers

import (
	"pokemon/m/v1/models"
	"sort"
)

func SortPokemonsBasedOnEditDistance(pokemonsWithEditDistance []models.PokemonsWithEditDistanceStruct) []models.PokemonsWithEditDistanceStruct {
	sort.Slice(pokemonsWithEditDistance, func(i, j int) bool {
		return pokemonsWithEditDistance[i].EditDistance < pokemonsWithEditDistance[j].EditDistance
	})
	return pokemonsWithEditDistance
}
