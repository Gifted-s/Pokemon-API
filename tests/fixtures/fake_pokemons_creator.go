package fixtures

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pokemon/m/v1/models"
)

func CreateFakePokemons() []models.Pokemon {
	fakeId1, _ := primitive.ObjectIDFromHex("62a705c3cfb07e0c463b315b")
	fakeId2, _ := primitive.ObjectIDFromHex("62a705c5cfb07e0c463b315c")
	fakeId3, _ := primitive.ObjectIDFromHex("62a705c5cfb07e0c463b315d")
	fakeId4, _ := primitive.ObjectIDFromHex("62a705c6cfb07e0c463b315e")
	fakeId5, _ := primitive.ObjectIDFromHex("62a705c8cfb07e0c463b3167")
	fakeId6, _ := primitive.ObjectIDFromHex("62a705c8cfb07e0c463b3168")
	fakeId7, _ := primitive.ObjectIDFromHex("62a705c8cfb07e0c463b3169")
	fakeId8, _ := primitive.ObjectIDFromHex("62a705c8cfb07e0c463b316a")
	fakeId9, _ := primitive.ObjectIDFromHex("62a705c8cfb07e0c463b316b")
	fakeId10, _ := primitive.ObjectIDFromHex("62a705c9cfb07e0c463b316c")
	fakeId11,_:= primitive.ObjectIDFromHex("62a705c9cfb07e0c463b316d")
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
			DefenceSpeed: 10,
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
			DefenceSpeed: 10,
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
			DefenceSpeed: 10,
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
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
		{
			ID:           fakeId5,
			Name:         "Galgas",
			Type1:        "Water",
			Type2:        "Grass",
			Total:        34,
			HP:           50,
			Attack:       456,
			Defense:      54,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
		{
			ID:           fakeId6,
			Name:         "Galgas",
			Type1:        "Grass",
			Type2:        "Grass",
			Total:        34,
			HP:           50,
			Attack:       456,
			Defense:      54,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
		{
			ID:           fakeId7,
			Name:         "Papeteer",
			Type1:        "Water",
			Type2:        "Grass",
			Total:        34,
			HP:           50,
			Attack:       456,
			Defense:      54,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},

		{
			ID:           fakeId8,
			Name:         "Thunder",
			Type1:        "Water",
			Type2:        "Grass",
			Total:        34,
			HP:           50,
			Attack:       456,
			Defense:      54,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},

		{
			ID:           fakeId9,
			Name:         "ThunderBolt",
			Type1:        "Water",
			Type2:        "Grass",
			Total:        34,
			HP:           50,
			Attack:       456,
			Defense:      54,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
		{
			ID:           fakeId10,
			Name:         "Kilosko",
			Type1:        "Bug",
			Type2:        "Grass",
			Total:        34,
			HP:           50,
			Attack:       456,
			Defense:      54,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},

		{
			ID:           fakeId11,
			Name:         "Jumbo",
			Type1:        "Flying",
			Type2:        "Grass",
			Total:        34,
			HP:           50,
			Attack:       456,
			Defense:      54,
			AttackSpeed:  600,
			DefenceSpeed: 10,
			Speed:        90,
			Generation:   1,
			Lengendary:   false,
		},
	}
	return fakeRandomPokemons
}
