package controller

import (
	"encoding/json"
	"net/http"
	"pokemon/m/v1/db"
	"pokemon/m/v1/internal/helpers"
	"pokemon/m/v1/models"
)

// GetPokemonsController sends an HTTP success response that contains pokemons that matches the request query params.
// It sends an error status and error message if an error was encountered.
func GetPokemonsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParams := r.URL.Query()
	// create a customized version of the query params for further consumption by Database
	customParams, err := helpers.CustomizeQueryParams(queryParams)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := models.ErrorResponseStruc{Status: http.StatusBadRequest, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}
    // Fetch pokemons from Database
	pokemons, err := db.GetPokeMons(customParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := models.ErrorResponseStruc{Status: http.StatusInternalServerError, ErrorMsg: err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	searchText := ""
	// check if a search query param was passed 
	if val, ok := customParams["search"]; ok {
		searchText = val[0]
	}
	// compute edit distance between pokemon names and the search text
	pokemonsWithEditDistance := helpers.ComputeLevenshteinDistance(pokemons, searchText)
	// sort pokemons based on edit distance
	sortedPokemonsBasedonEditDistance := helpers.SortPokemonsBasedOnEditDistance(pokemonsWithEditDistance)
	// create pokemon slice to be returned in response body
	page := "1"
	if v, ok := customParams["page"]; ok{
      page =v[0]
	}
	pokemonsSliceContructor := helpers.PokemonSliceConstructor(sortedPokemonsBasedonEditDistance, page)
	w.WriteHeader(http.StatusOK)
	// create success response body
	resp := models.GetPokemonsSuccessResponseStruc{Status: http.StatusOK, Pokemons: pokemonsSliceContructor}
	// return response to user in JSON format
	json.NewEncoder(w).Encode(resp)
}
