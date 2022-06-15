package helpers

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pokemon/m/v1/models"
	"testing"
)

func Test_Compute_Edit_Distance(t *testing.T) {
	assert := assert.New(t)
	searchWord := "Chaldeans"
	fakeId1, _ := primitive.ObjectIDFromHex("62a705c3cfb07e0c463b315b")
	fakeId2, _ := primitive.ObjectIDFromHex("62a705c5cfb07e0c463b315c")
	fakeId3, _ := primitive.ObjectIDFromHex("62a705c5cfb07e0c463b315d")
	fakeId4, _ := primitive.ObjectIDFromHex("62a705c6cfb07e0c463b315e")

	fakeRandomPokemons := []models.Pokemon{
		{
			ID:           fakeId1,
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
			ID:           fakeId2,
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
		{
			ID:           fakeId3,
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
			Lengendary:   false,
		},
		{
			ID:           fakeId4,
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
			Lengendary:   false,
		},
	}
	pokemonsWithEditDistance := ComputeLevenshteinDistance(fakeRandomPokemons, searchWord)

	// Test individual pokeman name edit distance
	pokemon1ExpectedEditDistance := 6
	assert.Equal(pokemon1ExpectedEditDistance, pokemonsWithEditDistance[0].EditDistance, "The edit distance between Chaldeans and Charmander is 6 ")

	pokemon2ExpectedEditDistance := 8
	assert.Equal(pokemon2ExpectedEditDistance, pokemonsWithEditDistance[1].EditDistance, "The edit distance between Chaldeans and Ceramic is 8 ")

	pokemon3ExpectedEditDistance := 0
	assert.Equal(pokemon3ExpectedEditDistance, pokemonsWithEditDistance[2].EditDistance, "The edit distance between Chaldeans and Chaldeans is 0 ")

	pokemon4ExpectedEditDistance := 6
	assert.Equal(pokemon4ExpectedEditDistance, pokemonsWithEditDistance[3].EditDistance, "The edit distance between Chaldeans and Cumimank is 6 ")
}
