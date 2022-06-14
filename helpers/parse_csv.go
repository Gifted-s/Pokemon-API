package helpers

import (
	"pokemon/m/v1/models"
	"strconv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"

)

func ParseCSV(rec []string) (*models.Pokemon, bool) {
	name := rec[1]
	type1 := rec[2]
	type2 := rec[3]
	total, _ := strconv.Atoi(rec[4])
	hp, _ := strconv.Atoi(rec[5])
	attack, _ := strconv.Atoi(rec[6])
	defense, _ := strconv.Atoi(rec[7])
	attackSpeed, _ := strconv.Atoi(rec[8])
	defenseSpeed, _ := strconv.Atoi(rec[9])
	speed, _ := strconv.Atoi(rec[10])
	generation, _ := strconv.Atoi(rec[11])
	legendary := rec[12]
	
	if legendary=="True" {
		return &models.Pokemon{}, false
	}
	if type1 == "Ghost" || type2 == "Ghost" {
		return &models.Pokemon{}, false
	}

	if type1 == "Steel" || type2 == "Steel" {
		hp *= 2
	}

	if type1 == "Fire" || type2 == "Fire" {
		tenPercentOfAttack := math.Round(float64(attack) * 0.1)
		attack -= int(tenPercentOfAttack)
	}

	if type1 == "Bug" || type2 == "Flying" || type1 == "Flying" || type2 == "Bug" {
		tenPercentOfAttackSpeed := math.Round(float64(attackSpeed) * 0.1)
		attackSpeed += int(tenPercentOfAttackSpeed)
	}

	if char := name[0:1]; char == "G" {
		nameRunes := []rune(name)
		for i := 1; i < len(nameRunes); i++ {
			defense += 5
		}
	}
	pokemon := &models.Pokemon{
		ID:           primitive.NewObjectID(),
		Name:         name,
		Type1:        type1,
		Type2:        type2,
		Total:        total,
		HP:           hp,
		Attack:       attack,
		Defense:      defense,
		AttackSpeed:  attackSpeed,
		DefenseSpeed: defenseSpeed,
		Speed:        speed,
		Generation:   generation,
		Lengendary:   false,// Note we don't want pokemons with legendary type
	}
	return pokemon, true

}
