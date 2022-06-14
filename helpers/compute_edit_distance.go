package helpers

import (
	"pokemon/m/v1/models"
	"strings"
)

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
