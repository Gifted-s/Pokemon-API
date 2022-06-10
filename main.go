package main
import (
	"github.com/joho/godotenv"
	"log"
)


func main() {
	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
}
