package helpers

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pokemon/m/v1/models"
	"reflect"
	"testing"
)

func TestConstructPokemonSlice(t *testing.T) {
	assert := assert.New(t)
	fakeId1, _ := primitive.ObjectIDFromHex("62a705c3cfb07e0c463b315b")
	fakeId2, _ := primitive.ObjectIDFromHex("62a705c5cfb07e0c463b315c")

	fakePokemanSliceWithEditDistance := []*models.PokemonsWithEditDistanceStruct{
		{EditDistance: 6,
			Pokemon: models.Pokemon{
				ID: fakeId1,
				Name:         "Charmander",
				Type1:        "Fire",
				Type2:        "Grass",
				Total:        100,
				HP:           50,
				Attack:       400,
				Defense:      200,
				AttackSpeed:  600,
				DefenseSpeed: 10,
				Speed:        90,
				Generation:   1,
				Lengendary:   false},
		},
		{EditDistance: 8,
			Pokemon: models.Pokemon{
				ID: fakeId2,
				Name:         "Ceramic",
				Type1:        "Fire",
				Type2:        "Grass",
				Total:        100,
				HP:           50,
				Attack:       456,
				Defense:      54,
				AttackSpeed:  600,
				DefenseSpeed: 10,
				Speed:        90,
				Generation:   1,
				Lengendary:   false},
		},
	}


	expectedPokemonSlice := []*models.Pokemon{
		{
			
				ID: fakeId1,
				Name:         "Charmander",
				Type1:        "Fire",
				Type2:        "Grass",
				Total:        100,
				HP:           50,
				Attack:       400,
				Defense:      200,
				AttackSpeed:  600,
				DefenseSpeed: 10,
				Speed:        90,
				Generation:   1,
				Lengendary:   false,
		},
		{
				ID: fakeId2,
				Name:         "Ceramic",
				Type1:        "Fire",
				Type2:        "Grass",
				Total:        100,
				HP:           50,
				Attack:       456,
				Defense:      54,
				AttackSpeed:  600,
				DefenseSpeed: 10,
				Speed:        90,
				Generation:   1,
				Lengendary:   false,
		},
	}
    sortedPokemons := PokemonSliceConstructor(fakePokemanSliceWithEditDistance)
	slicesAreEqual := reflect.DeepEqual(expectedPokemonSlice, sortedPokemons)
	assert.Equal(true, slicesAreEqual, "The two slices must be equal")

}
