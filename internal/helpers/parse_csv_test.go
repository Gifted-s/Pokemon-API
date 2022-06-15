package helpers

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Exclude_Legendary_Pokemon(t *testing.T) {
	assert := assert.New(t)
	lengendary := "True"
	csvRow := []string{"1","Bulbasaur","Grass","Poison","318","45","49","49","65","65","45","1",lengendary}
	_, parsed := ParseCSV(csvRow)
	assert.Equal(false, parsed, "parsed should be false since lengendary is true")

}



func Test_Exclude_Ghost_Type(t *testing.T) {
	assert := assert.New(t)
	type1 := "Ghost"
	csvRow := []string{"1","Bulbasaur",type1,"Poison","318","45","49","49","65","65","45","1","False"}
	_, parsed := ParseCSV(csvRow)
	assert.Equal(false, parsed, "parsed should be false since Type1 is Ghost")

	type2 := "Ghost"
	csvRow2 := []string{"2","Bulbasaur","Grass",type2,"318","45","49","49","65","65","45","1","False"}
	_, parsed2 := ParseCSV(csvRow2)
	assert.Equal(false, parsed2, "parsed should be false since Type2 is Ghost")

}



func Test_Double_Hp_For_Type_Steel(t *testing.T) {
	assert := assert.New(t)
	hp:="45"
	csvRow := []string{"1","Bulbasaur","Steel","Poison","318",hp,"49","49","65","65","45","1","False"}
	pokemon, parsed := ParseCSV(csvRow)
	assert.Equal(90, pokemon.HP, "HP should be doubled since Type1 is Steel")
	assert.Equal(true, parsed, "parsed should be true since since Type1 and Type2 are valid types")

	hp2:="50"
	csvRow2 := []string{"2","Bulbasaur","Grass","Steel","318",hp2,"49","49","65","65","45","1","False"}
	pokemon2, parsed2 := ParseCSV(csvRow2)
	assert.Equal(100, pokemon2.HP, "HP should be doubled since Type2 is Steel")
	assert.Equal(true, parsed2, "parsed should be true since since Type1 and Type2 are valid types")

}

func Test_Reduce_Atttack_By_10_Percent_For_Type_Fire(t *testing.T) {
	assert := assert.New(t)
	attack:="50"
	csvRow := []string{"1","Bulbasaur","Fire","Poison","318","45",attack,"49","65","65","45","1","False"}
	pokemon, parsed := ParseCSV(csvRow)
	assert.Equal(45, pokemon.Attack, "Attack should be reduced by 10 Percent since Type1 is Fire")
	assert.Equal(true, parsed, "parsed should be true since since Type1 and Type2 are valid types")

	attack2:="60"
	csvRow2 := []string{"2","Bulbasaur","Grass","Fire","318","50",attack2,"49","65","65","45","1","False"}
	pokemon2, parsed2 := ParseCSV(csvRow2)
	assert.Equal(54, pokemon2.Attack, "Attack should be reduced by 10 Percent since Type2 is Fire")
	assert.Equal(true, parsed2, "parsed should be true since since Type1 and Type2 are valid types")

}


func Test_Increase_Attack_Speed_By_10_Percent_For_Type_Bug(t *testing.T) {
	assert := assert.New(t)
	attackSpeed:= "60";

	csvRow := []string{"1","Bulbasaur","Bug","Poison","318","45","50","49",attackSpeed,"65","45","1","False"}
	pokemon, parsed := ParseCSV(csvRow)
	assert.Equal(66, pokemon.AttackSpeed, "Attack Speed should be increased by 10 Percent since Type1 is Bug")
	assert.Equal(true, parsed, "parsed should be true since since Type1 and Type2 are valid types")

	attackSpeed2:= "70";

	csvRow2 := []string{"2","Bulbasaur","Grass","Bug","318","50","60","49",attackSpeed2,"65","45","1","False"}
	pokemon2, parsed2 := ParseCSV(csvRow2)
	assert.Equal(77, pokemon2.AttackSpeed, "Attack Speed should be increased by 10 Percent since Type2 is Bug")
	assert.Equal(true, parsed2, "parsed should be true since since Type1 and Type2 are valid types")

}


func Test_Increase_Attack_Speed_By_10_Percent_For_Type_Flying(t *testing.T) {
	assert := assert.New(t)
	attackSpeed:= "60";

	csvRow := []string{"1","Bulbasaur","Flying","Poison","318","45","50","49",attackSpeed,"65","45","1","False"}
	pokemon, parsed := ParseCSV(csvRow)
	assert.Equal(66, pokemon.AttackSpeed, "Attack Speed should be increased by 10 Percent since Type1 is Bug")
	assert.Equal(true, parsed, "parsed should be true since since Type1 and Type2 are valid types")

	attackSpeed2:= "70";

	csvRow2 := []string{"2","Bulbasaur","Grass","Flying","318","50","60","49",attackSpeed2,"65","45","1","False"}
	pokemon2, parsed2 := ParseCSV(csvRow2)
	assert.Equal(77, pokemon2.AttackSpeed, "Attack Speed should be increased by 10 Percent since Type2 is Bug")
	assert.Equal(true, parsed2, "parsed should be true since since Type1 and Type2 are valid types")

}


func Test_Add_5_For_Each_Char_If_Name_Start_With_G_Except_G(t *testing.T) {
	assert := assert.New(t)
	
	defense:= "60";

	csvRow := []string{"1","Gloom","Flying","Poison","318","45","60",defense,"60","65","45","1","False"}
	pokemon, parsed := ParseCSV(csvRow)
	assert.Equal(80, pokemon.Defense, "Defense to be increased by 20 since we add 5 for each char except G")
	assert.Equal(true, parsed, "parsed should be true since since Type1 and Type2 are valid types")
}