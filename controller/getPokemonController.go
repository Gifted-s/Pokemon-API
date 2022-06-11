package controller

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"pokemon/m/v1/db"
	"pokemon/m/v1/helpers"
)

func GetPokemonsController(w http.ResponseWriter, r *http.Request) {
	queryParams :=r.URL.Query()
	customParams:= helpers.CustomizeQueryParams(queryParams)
	 db.GetPokeMons(customParams)
	
	resp := map[string]string{"Response": "response"}
	json.NewEncoder(w).Encode(resp)
}