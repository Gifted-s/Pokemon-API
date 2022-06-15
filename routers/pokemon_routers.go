package routers

import (
	"pokemon/m/v1/controller"
	"github.com/gorilla/mux"
)

// PokemonRouter is responsible for handling request for various paths, it return a mux router type 
func PokemonRouter(r *mux.Router) *mux.Router {
	r.HandleFunc("/pokemon", controller.GetPokemonsController).Methods("GET")
	return r
}
