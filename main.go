package main

import (
	//"github.com/eaacisternas/pokeBack/database"

	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	//"github.com/eaacisternas/pokeBack/handlers"
)

type Response struct {
	Pokemon []Pokemon `json:"pokemon_entries"`
}
type Pokemon struct {
	EntryNo int `json:"entry_number"`
}

type NewPokemon struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Image  Images  `json:"sprite"`
	Weight int     `json:"weight"`
	Types  []Types `json:"types"`
}
type Types struct {
	Type string `json:"type{name}"`
}
type Images struct {
	Url string `json:"front_default"`
}

func main() {
	// if bd.ValConnection() == 0 {
	// 	log.Fatal("No se pudo conectar a la bd")
	// 	return
	// }
	//handlers.Manejadores()
	leerPokeApi("kanto")
}

func leerPokeApi(region string) {
	//la region sera modificable
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/" + region)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var responseObject Response
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
		var pokemonObject NewPokemon
		json.Unmarshal(responseData, &pokemonObject)
		println(string(pokemonObject.Name))
		println(int(pokemonObject.Id))
		println(int(pokemonObject.Weight))
		// println(pokemonObject.Types)
	}
}
