package bd

import (
	"context"
	"log"

	// "github.com/eaacisternas/pokeBack/middlewears"
	// "github.com/eaacisternas/pokeBack/routers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la base de datos*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://eduardoAguirre:Abril201118@cluster0.ep7djah.mongodb.net/?retryWrites=true&w=majority")

/*ConerctarBD es la función que me permite conectar la BD*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa con la base de datos")
	return client
}

/*valConnectiones hace el ping a la Bd*/
func ValConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
