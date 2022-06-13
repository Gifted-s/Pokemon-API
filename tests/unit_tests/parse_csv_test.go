package parse_csv_test

import (
	"github.com/stretchr/testify/assert"
	"pokemon/m/v1/helpers"
	"testing"
)

func CsvParserTest(t *testing.T) {
	// models.Pokemon, bool
	assert := assert.New(t)
	csvRow := []string{`
	#,Name,Type 1,Type 2,Total,HP,Attack,Defense,Sp. Atk,Sp. Def,Speed,Generation,Legendary`,
		`1,Bulbasaur,Grass,Poison,318,45,49,49,65,65,45,1,True`}
	_, parsed := helpers.ParseCSV(csvRow)
	assert.Equal(parsed, false, "parsed should be false since lengendary is true")

}
