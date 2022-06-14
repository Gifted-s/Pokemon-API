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
		w.WriteHeader(http.StatusBadRequest)
		resp := models.ErrorResponseStruc{Status: http.StatusBadRequest, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	pokemons, err := db.GetPokeMons(customParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := models.ErrorResponseStruc{Status: http.StatusInternalServerError, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	searchText := ""
	if val, ok := customParams["search"]; ok {
		searchText = val[0]
	}
	pokemonsWithEditDistance := helpers.ComputeLevenshteinDistance(pokemons, searchText)
	sortedPokemonsBasedonEditDistance := helpers.SortPokemonsBasedOnEditDistance(pokemonsWithEditDistance)
	pokemonsSliceContructor := helpers.PokemonSliceConstructor(sortedPokemonsBasedonEditDistance)
	w.WriteHeader(http.StatusOK)
	resp := models.GetPokemonsSuccessResponseStruc{Status: http.StatusOK, Pokemons: pokemonsSliceContructor}
	json.NewEncoder(w).Encode(resp)
}
