package controller

import (
	"encoding/json"
	"net/http"
	"pokemon/m/v1/db"
	"pokemon/m/v1/helpers"
	"pokemon/m/v1/models"
)

func GetPokemonsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParams := r.URL.Query()
	customParams, err := helpers.CustomizeQueryParams(queryParams)
	if err != nil {
		resp := models.ErrorResponseStruc{Status: 400, ErrorMsg: err.Error()}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			panic(err)
		}
	} else {
		pokemons, err := db.GetPokeMons(customParams)
		if err != nil {
			resp := models.ErrorResponseStruc{Status: 400, ErrorMsg: err.Error()}
			err := json.NewEncoder(w).Encode(resp)
			if err != nil {
				panic(err)
			}
		} else {
			searchText := ""
			if val, ok := customParams["search"]; ok {
				searchText = val[0]
			}
			pokemonsWithEditDistance := helpers.ComputeLevenshteinDistance(pokemons, searchText)
			sortedPokemonsBasedonEditDistance := helpers.SortPokemonsBasedOnEditDistance(pokemonsWithEditDistance)
			resp := models.GetPokemonsSuccessResponseStruc{Status: 200, Pokemons: sortedPokemonsBasedonEditDistance}
			err := json.NewEncoder(w).Encode(resp)
			if err != nil {
				panic(err)
			}
		}
	}

}
