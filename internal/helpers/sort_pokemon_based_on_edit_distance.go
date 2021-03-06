package helpers

import (
	"pokemon/m/v1/models"
	"sort"
)
// SortPokemonsBasedOnEditDistance takes in pokemons with their edit distance and sort them based on the edit distance from smallest to largest
func SortPokemonsBasedOnEditDistance(pokemonsWithEditDistance []*models.PokemonsWithEditDistanceStruct) []*models.PokemonsWithEditDistanceStruct {
	sort.Slice(pokemonsWithEditDistance, func(i, j int) bool {
		return pokemonsWithEditDistance[i].EditDistance < pokemonsWithEditDistance[j].EditDistance
	})
	return pokemonsWithEditDistance
}
