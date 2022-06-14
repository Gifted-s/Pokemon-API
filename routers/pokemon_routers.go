package routers

import (
	"pokemon/m/v1/controller"
	"github.com/gorilla/mux"
)


func PokemonRouter(r *mux.Router) *mux.Router {
	r.HandleFunc("/pokemon", controller.GetPokemonsController).Methods("GET")
	return r
}
