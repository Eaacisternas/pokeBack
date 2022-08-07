package main

import (
	"log"

	"github.com/eaacisternas/pokeBack/handlers"
)

func main() {
	if bd.ValConnection() == 0 {
		log.Fatal("No se pudo conectar a la bd")
		return
	}
	handlers.Manejadores()
}
