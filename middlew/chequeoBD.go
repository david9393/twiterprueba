package middlew

import (
	"net/http"

	"github.com/david9393/twiterprueba/bd"
)

/* es el middlew que me permite conocer el esta de la base de datos */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion fallida con base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
