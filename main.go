package main
import (
	"github.com/joho/godotenv"
	"pokemon/m/v1/server"
	"log"
)


func main() {
	err := godotenv.Load(".env")

	if err != nil {
	  log.Panic("Error loading .env file")
	}
    created, err:= server.SetupServer()
	if !created{
		log.Panic(err)
	}else{
     log.Print("Server Created")
	}
}
