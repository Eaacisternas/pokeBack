package bd

import (
	"context"
	"time"

	"github.com/eaacisternas/pokeBack/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertarPokemon funcion con la cual se insertaran los pokemon en la base de datos*/
func InsertarPokemon(p models.Pokemon) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("pokeBack")
	col := db.Collection("PokemonCollection")
	result, err := col.InsertOne(ctx, p)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
