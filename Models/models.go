package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Pokemon is struct that represents the model of a pokemon. 
// Pokemon can be marshalled to BSON to be stored in Database or marshed to JSON to be sent through HTTP 
type Pokemon struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Type1        string             `json:"type1" bson:"type1,omitempty"`
	Type2        string             `json:"type2" bson:"type2,omitempty"`
	Total        int                `json:"total" bson:"total,omitempty"`
	HP           int                `json:"hp" bson:"hp,omitempty"`
	Attack       int                `json:"attack" bson:"attack,omitempty"`
	Defense      int                `json:"defense" bson:"defense,omitempty"`
	AttackSpeed  int                `json:"attackSpeed" bson:"attackSpeed,omitempty"`
	DefenseSpeed int                `json:"defenseSpeed" bson:"defenseSpeed,omitempty"`
	Speed        int                `json:"speed" bson:"speed,omitempty"`
	Generation   int                `json:"generation" bson:"generation,omitempty"`
	Lengendary   bool               `json:"lengendary" bson:"lengendary"`
}

// GetPokemonsSuccessResponseStruc is used to structure success response body for a get pokemon operation
type GetPokemonsSuccessResponseStruc struct {
	Status   int       `json:"status,omitempty"`
	Pokemons []*Pokemon `json:"pokemons" bson:"pokemons"`
}

// ErrorResponseStruc is used to structure error response body.
// ErrorResponseStruc can be marshalled to JSON
type ErrorResponseStruc struct {
	Status   int    `json:"status,omitempty"`
	ErrorMsg string `json:"error"`
}

// Configuration represent the model for app config
type Configuration struct {
	Db_Config Db_Config_Struct
}

// Db_Config_Struct is used to model the Database Config, it is includes various db collection names, db names and also fetch size for each query
type Db_Config_Struct struct {
	Pokemon_Collection_Name      string
	Pokemon_Collection_Name_Test string
	Db_Name_Test                 string
	Db_Name                      string
	Fetch_Limit                  int
}

// Collections is used to represent all the various collections we have in the database
type Collections struct {
	PokeMons *mongo.Collection
}

// GetPokemonQueryFormatterStruct is used to model the format the query to fetch pokemon from database
type GetPokemonQueryFormatterStruct struct {
	QueryCondition []bson.M
	Options        *options.FindOptions
}

// PokemonsWithEditDistanceStruct is used to model a pokemon and its edit distance from the search text  
type PokemonsWithEditDistanceStruct struct {
	EditDistance int
	Pokemon
}
