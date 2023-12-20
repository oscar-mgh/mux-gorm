package handlers

import (
	"clase_5_mux_gorm/database"
	"clase_5_mux_gorm/dto"
	"clase_5_mux_gorm/modelos"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

func Categoria_get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	datos := modelos.Categorias{}
	database.Database.Order("id desc").Find(&datos)
	//database.Database.Find(&datos)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(datos)
}
func Categoria_get_con_parametro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	//id, _ = strconv.Atoi(vars["id"])
	datos := modelos.Categoria{}
	if err := database.Database.First(&datos, vars["id"]); err.Error != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(respuesta)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(datos)
	}

}
func Categoria_post(w http.ResponseWriter, r *http.Request) {
	var categoria dto.CategoriaDto
	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	datos := modelos.Categoria{Nombre: categoria.Nombre, Slug: slug.Make(categoria.Nombre)}
	database.Database.Save(&datos)
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Se creó el registro exitosamente",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
func Categoria_put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var categoria dto.CategoriaDto

	if err := json.NewDecoder(r.Body).Decode(&categoria); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	datos := modelos.Categoria{}
	if err := database.Database.First(&datos, id); err.Error != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(respuesta)
		return
	} else {
		datos.Nombre = categoria.Nombre
		datos.Slug = slug.Make(categoria.Nombre)
		database.Database.Save(&datos)
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "Se modificó el registro exitosamente",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(respuesta)
	}
}
func Categoria_delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	datos := modelos.Categoria{}
	if err := database.Database.First(&datos, id); err.Error != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Recurso no disponible",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(respuesta)
		return
	} else {
		database.Database.Delete(&datos)
		respuesta := map[string]string{
			"estado":  "ok",
			"mensaje": "Se eliminó el registro exitosamente",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(respuesta)
	}
}
