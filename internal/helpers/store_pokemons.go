package helpers

import (
	"context"
	"encoding/csv"
	"pokemon/m/v1/db"
	"fmt"
	"io"
	"log"
	"os"
)

// StorePokemonInDB is responsible for pulling each row from the CSV file, converting them to a pokemon model and inserting the pokemon into the Database
func StorePokemonInDB() {
	pokemonCollection := db.GetDBCollections().PokeMons
	// open file
	f, err := os.Open("./assets/Data/pokemon.csv")
	if err != nil {
		log.Fatal(err)
	}
	// close the file at the end of the program
	defer f.Close()
	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	// Read the first row to prevent inserting Label row in the loop
	csvReader.Read()
	// Insert remaining row
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		customPokeMon, addToDB := ParseCSV(rec)
		if addToDB {
			result, err := pokemonCollection.InsertOne(context.TODO(), customPokeMon)
			if err != nil {
				log.Panic(err)
				return
			}
			 fmt.Print(result)
		}

	}

}
