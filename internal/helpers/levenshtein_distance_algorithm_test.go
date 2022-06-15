package helpers


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Levenshtein_Distance_Algorithm(t *testing.T) {
	assert := assert.New(t)
	editDistance := MinDistance("bola", "bulbasaur")
	assert.Equal(6, editDistance, "edit distance should be 6 since it takes six operations to convert bola to  bulbasaur")
}


func Test_Levenshtein_Distance_Algorithm_With_Two_Identical_Strings(t *testing.T) {
	assert := assert.New(t)
	editDistance := MinDistance("Metapod", "Metapod")
	assert.Equal(0, editDistance, "edit distance should be 0 since it takes 0 operations to convert Metapod to Metapod")
}

func Test_Levenshtein_Distance_Algorithm_With_Empty_String1(t *testing.T) {
	assert := assert.New(t)
	editDistance := MinDistance("", "bulbasaur")
	assert.Equal(9, editDistance, "edit distance should be 9 since it takes 9 insertions to convert empty string to  bulbasaur")
}


func Test_Levenshtein_Distance_Algorithm_With_Empty_String2(t *testing.T) {
	assert := assert.New(t)
	editDistance := MinDistance("bulbasaur", "")
	assert.Equal(9, editDistance, "edit distance should be 9 since it takes 9 deletions to convert bulbasaur to an empty string")
}


func Test_Levenshtein_Distance_Algorithm_With_Two_Empty_Strings(t *testing.T) {
	assert := assert.New(t)
	editDistance := MinDistance("", "")
	assert.Equal(0, editDistance, "edit distance should be 0 since it takes 0 operations to convert an empty string to an empty string")
}