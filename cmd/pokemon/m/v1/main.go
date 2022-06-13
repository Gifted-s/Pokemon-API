package main

import (
	"log"
	//"pokemon/m/v1/helpers"
	"pokemon/m/v1/server"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load(".env")

	if err != nil {
	  log.Panic("Error loading .env file")
	}
	//helpers.StorePokemonInDB()
    created, err:= server.SetupServer()
	if !created{
		log.Panic(err)
	}else{
     log.Print("Server Created")
	}
}
