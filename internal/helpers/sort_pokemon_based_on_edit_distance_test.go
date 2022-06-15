package helpers

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pokemon/m/v1/models"
	"testing"
)

func Test_Sort_Pokemon_Based_On_Edit_Distance(t *testing.T) {
	assert := assert.New(t)
	fakeId1, _ := primitive.ObjectIDFromHex("62a705c3cfb07e0c463b315b")
	fakeId2, _ := primitive.ObjectIDFromHex("62a705c5cfb07e0c463b315c")
	fakeId3, _ := primitive.ObjectIDFromHex("62a705c5cfb07e0c463b315d")
	fakeId4, _ := primitive.ObjectIDFromHex("62a705c6cfb07e0c463b315e")


	

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
		{EditDistance: 0,
			Pokemon: models.Pokemon{
				ID: fakeId3,
				Name:         "Chaldeans",
				Type1:        "Water",
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
		{EditDistance: 6,
			Pokemon: models.Pokemon{
				ID: fakeId4,
				Name:         "Cumimank",
				Type1:        "Water",
				Type2:        "Grass",
				Total:        34,
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

    sortedPokemons := SortPokemonsBasedOnEditDistance(fakePokemanSliceWithEditDistance)
	assert.Equal(sortedPokemons[0].ID, fakeId3, "The pokemon with fakeId3 should come first since it has edit distance 0")
	assert.Equal(sortedPokemons[1].ID, fakeId1, "The pokemon with fakeId1 should come second since it has edit distance 6")
	assert.Equal(sortedPokemons[2].ID, fakeId4, "The pokemon with fakeId4 should come third since it has edit distance 6")
	assert.Equal(sortedPokemons[3].ID, fakeId2, "The pokemon with fakeId2 should come fourth since it has edit distance 8")
}
