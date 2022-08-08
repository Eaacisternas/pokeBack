package middleware

import (
	"net/http"

	bd "github.com/eaacisternas/pokeBack/database"
)

/*ValBD funcion hecha para funcionar como middleware para conocer el estado de la base de datos*/
func ValBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ValConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
