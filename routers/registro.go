package routers

import (
	"encoding/json"
	"net/http"

	"github.com/david9393/twiterprueba/bd"
	"github.com/david9393/twiterprueba/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email no valido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "El password debe tener almenos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuarion con ese email", 400)
		return
	}
	_, status, err := bd.InsertarRegistro(t)
	if err != nil {
		http.Error(w, "Error al insertar el usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se a logrado insertar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
