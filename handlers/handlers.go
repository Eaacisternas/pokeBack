package handlers

import (
	"log"
	"net/http"
	"os"

	middleware "github.com/eaacisternas/pokeBack/middlewears"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el handler y agrero un listener al server*/
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/procesarListadoPokemon", middleware.ValBD(routers.procesarListadoPokemon)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
