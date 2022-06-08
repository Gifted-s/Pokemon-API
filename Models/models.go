package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pokemon struct {
	MongoDBID  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Type1      string             `json:"type1" bson:"type1,omitempty"`
	Type2      string             `json:"type2" bson:"type2,omitempty"`
	Total      int                `json:"total" bson:"total,omitempty"`
	HP         int                `json:"hp" bson:"hp,omitempty"`
	Attack     int                `json:"attack" bson:"attack ,omitempty"`
	Defence    int                `json:"defence" bson:"defence ,omitempty"`
	SPAttack   int                `json:"sPattack" bson:"sPattack ,omitempty"`
	SPDefence  int                `json:"sPdefence" bson:"sPdefence ,omitempty"`
	Speed      int                `json:"speed" bson:"speed ,omitempty"`
	Generation int                `json:"generation" bson:"generation ,omitempty"`
	Lengendary int                `json:"lengendary" bson:"lengendary ,omitempty"`
}

