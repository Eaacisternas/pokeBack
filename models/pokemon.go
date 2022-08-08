package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pokemon struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Orden  int                `bson:"order" json:"order,omitempty"`
	Nombre string             `bson:"name" json:"name,omitempty"`
	Peso   int                `bson:"weight" json:"weight,omitempty"`
	Altura int                `bson:"height" json:"height,omitempty"`
}

// type NewPokemon struct {
// 	Id        int         `json:"id"`
// 	Name      string      `json:"name"`
// 	Abilities []Abilities `json:"abilities"`
// 	Weight    int         `json:"weight"`
// }
// type Abilities struct {
// 	Slot int `json:"slot"`
// }

// type Types struct {
// 	Type string `json:"type{name}"`
// }
// type Images struct {
// 	Url string `json:"front_default"`
// }
