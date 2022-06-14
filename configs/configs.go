package configs

import (
"pokemon/m/v1/models")

var Config = models.Configuration {

	Db_Config :models.Db_Config_Struct{
	Pokemon_Collection_Name:"pokemons",
	Pokemon_Collection_Name_Test:"pokemons_test" ,
	Db_Name:"pokemon_db",
	Db_Name_Test: "pokemon_db_test",
	Fetch_Limit: 10,
    },

}

