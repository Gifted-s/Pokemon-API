package helpers

import (
	"pokemon/m/v1/models"
)

// PokemonSliceConstructor takes pokemons, and returns the pokemons but removes the edit distance field
func PokemonSliceConstructor(pokemonsWithEditDistance []*models.PokemonsWithEditDistanceStruct) []*models.Pokemon {
	pokemonSlice := []*models.Pokemon{}
	for _, p := range pokemonsWithEditDistance {
		pokemonSlice = append(pokemonSlice,&models.Pokemon{
			ID:           p.ID,
			Name:         p.Name,
			Type1:        p.Type1,
			Type2:        p.Type2,
			Total:        p.Total,
			HP:           p.HP,
			Attack:       p.Attack,
			Defense:      p.Defense,
			AttackSpeed:  p.AttackSpeed,
			DefenseSpeed: p.DefenseSpeed,
			Speed:        p.Speed,
			Generation:   p.Generation,
			Lengendary:   false,
		})
	}
	return pokemonSlice
}
