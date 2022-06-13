package configs

import (
"pokemon/m/v1/models")

var Config = models.Configuration {

	Db_Config :models.Db_Config_Struct{
	Pokemon_Collection_Name:"pokemons",
	Db_Name:"pokemon_db",
	Fetch_Limit: 10,
    },

}

