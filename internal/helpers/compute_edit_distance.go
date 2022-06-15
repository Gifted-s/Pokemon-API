package helpers

import (
	"pokemon/m/v1/models"
	"strings"
)

// ComputeLevenshteinDistance is responsible for computing the minimum edit distance of each pokemon name from the search text.
// It returns a pokemon slice where each pokemon contains the edit distance and the pokemon fields
func ComputeLevenshteinDistance(pokemons []models.Pokemon, searchWord string) []*models.PokemonsWithEditDistanceStruct {
	pokemonsWithWithEditDistanceSlice := []*models.PokemonsWithEditDistanceStruct{}
	for _, pokemon := range pokemons {
		editDistance := MinDistance(strings.ToLower(searchWord), strings.ToLower(pokemon.Name))
		pokemonsWithWithEditDistanceSlice = append(pokemonsWithWithEditDistanceSlice, &models.PokemonsWithEditDistanceStruct{
			EditDistance: editDistance,
			Pokemon:      pokemon,
		})
	}
	return pokemonsWithWithEditDistanceSlice
}
