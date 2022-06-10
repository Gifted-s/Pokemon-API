package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Pokemon struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Type1        string             `json:"type1" bson:"type1,omitempty"`
	Type2        string             `json:"type2" bson:"type2,omitempty"`
	Total        int                `json:"total" bson:"total,omitempty"`
	HP           int                `json:"hp" bson:"hp,omitempty"`
	Attack       int                `json:"attack" bson:"attack ,omitempty"`
	Defence      int                `json:"defence" bson:"defence ,omitempty"`
	AttackSpeed  int                `json:"attackSpeed" bson:"attackSpeed ,omitempty"`
	DefenceSpeed int                `json:"defenceSpeed" bson:"defenceSpeed ,omitempty"`
	Speed        int                `json:"speed" bson:"speed ,omitempty"`
	Generation   int                `json:"generation" bson:"generation ,omitempty"`
	Lengendary   bool               `json:"lengendary" bson:"lengendary"`
}

type Configuration struct {
	Db_Config Db_Config_Struct
}

type Db_Config_Struct struct {
	Pokemon_Collection_Name string
	Db_Name                 string
}

type Collections struct {
	PokeMons *mongo.Collection
}
