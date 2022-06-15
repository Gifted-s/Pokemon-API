package helpers

import (
	"pokemon/m/v1/configs"
	"pokemon/m/v1/models"
	"strconv"
)

//  GetPage returns a part of the pokemons slice based on page number
func GetPage(pokemons []*models.Pokemon, pageInStringFormat string) []*models.Pokemon {
	page := 0
	page, _ = strconv.Atoi(pageInStringFormat)
	limit := configs.Config.Db_Config.Fetch_Limit
	if limit != 0 {
		if page == 0 {
			page = 1
		}
	}
	offset := (page - 1) * (limit + 1)
	end := offset + limit
	// If offset is larger than length of pokemon then return empty slice
	if offset > len(pokemons){
      return []*models.Pokemon{}
	}else if end > len(pokemons){
		return pokemons[offset:]
	}
	return pokemons[offset:end]
}
