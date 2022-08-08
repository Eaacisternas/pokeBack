package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	models "github.com/eaacisternas/pokeBack/models"
)

/*GuardarListado es la funcion para guardar la lista de json en la base de datos*/
func GuardarListado() {
	// var pokemonRecibidos models.Pokemon
	pokemonRecibidos := leerPokeApi("kanto")
	fmt.Println(pokemonRecibidos)

}

/*leerPokeApi es la funcion con la cual traemos los datos de los pokemon*/
func leerPokeApi(region string) models.Pokemon {
	//la region sera modificable

	var pokemonObject models.Pokemon
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/" + region)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var responseObject models.Region
	json.Unmarshal(responseData, &responseObject)
	for i := 1; i <= len(responseObject.Pokemon); i++ {
		response, err = http.Get("http://pokeapi.co/api/v2/pokemon/" + fmt.Sprint(i))
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err = io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		json.Unmarshal(responseData, &pokemonObject)
	}
	return pokemonObject
}
