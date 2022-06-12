package controller

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"pokemon/m/v1/db"
	"pokemon/m/v1/helpers"
	"pokemon/m/v1/models"
)

func GetPokemonsController(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	customParams := helpers.CustomizeQueryParams(queryParams)
	pokemons, err := db.GetPokeMons(customParams)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		resp := models.ErrorResponseStruc{Status: 400, Error: err}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			panic(err)
		}
	} else {
		resp := models.GetPokemonsSuccessResponseStruc{Status: 200, Pokemons: pokemons}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			panic(err)
		}
	}
}
