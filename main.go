package main

import (
	"log"

	bd "github.com/eaacisternas/pokeBack/database"
	"github.com/eaacisternas/pokeBack/handlers"
)

// type Response struct {
// 	Name    string    `json:"name"`
// 	Pokemon []Pokemon `json:"pokemon_entries"`
// }

// type Pokemon struct {
// 	EntryNo int `json:"entry_number"`
// }

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

func main() {
	if bd.ValConnection() == 0 {
		log.Fatal("No se pudo conectar a la bd")
		return
	}
	handlers.Manejadores()
	// fmt.Println("Conexion exitosa a la base de datos")
	// leerPokeApi("kanto")
}

// func leerPokeApi(region string) {
// 	//la region sera modificable
// 	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/" + region)

// 	if err != nil {
// 		fmt.Print(err.Error())
// 		os.Exit(1)
// 	}
// 	responseData, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	var responseObject Response
// 	json.Unmarshal(responseData, &responseObject)
// 	fmt.Println(responseObject.Name)
// 	fmt.Println(responseObject.Pokemon)
// 	for i := 1; i <= len(responseObject.Pokemon); i++ {
// 		response, err = http.Get("http://pokeapi.co/api/v2/pokemon/" + fmt.Sprint(i))
// 		if err != nil {
// 			fmt.Print(err.Error())
// 			os.Exit(1)
// 		}
// 		responseData, err = io.ReadAll(response.Body)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		var pokemonObject NewPokemon
// 		json.Unmarshal(responseData, &pokemonObject)
// 		println(pokemonObject.Abilities)
// 		// println(pokemonObject.Types)
// 	}
// }
