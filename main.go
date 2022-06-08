package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

func main() {
    // open file
    f, err := os.Open("./Data/pokemon.csv")
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)
    for {
        rec, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        // do something with read line
        fmt.Printf("%+v\n", reflect.TypeOf(rec))
		fmt.Println(rec[1], rec[2], rec[3],rec[4], rec[5], rec[6], rec[7], rec[8], rec[9], rec[10], rec[11], reflect.TypeOf(rec[12]))
    }
}