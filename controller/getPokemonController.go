package controller

import (
	"net/http"
	"encoding/json"
)

func GetPokemonsController(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"Response": "response"}
	json.NewEncoder(w).Encode(resp)
	
}