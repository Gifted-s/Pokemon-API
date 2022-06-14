package server

import (
	"log"
	"net/http"
	"pokemon/m/v1/db"
	"pokemon/m/v1/routers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupServer() (bool, error) {
	r := mux.NewRouter()
	http.Handle("/", r)
	routers.PokemonRouter(r)
	db.ConnectMongoDB()
	log.Println("Starting Server on Port 8080")
	err := http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r))
	if err != nil {
		return false, err
	}
	return true, nil
}
