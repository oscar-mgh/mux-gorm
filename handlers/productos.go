package handlers

import (
	"clase_5_mux_gorm/database"
	"clase_5_mux_gorm/dto"
	"clase_5_mux_gorm/modelos"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
)

func Productos_get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	datos := modelos.Productos{}
	database.Database.Order("id desc").Preload("Categoria").Find(&datos)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(datos)
}
func Productos_get_con_parametro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	//id, _ = strconv.Atoi(vars["id"])
	datos := modelos.Producto{}
	if err := database.Database.Preload("Categoria").First(&datos, vars["id"]); err.Error != nil {
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
func Productos_post(w http.ResponseWriter, r *http.Request) {
	var producto dto.ProductoDto
	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	datos := modelos.Producto{Nombre: producto.Nombre, Slug: slug.Make(producto.Nombre), Precio: producto.Precio, Stock: producto.Stock, Descripcion: producto.Descripcion, CategoriaID: producto.CategoriaID, Fecha: time.Now()}
	database.Database.Save(&datos)
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Se creó el registro exitosamente",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
func Productos_put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var producto dto.ProductoDto

	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	datos := modelos.Producto{}
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
		datos.Nombre = producto.Nombre
		datos.Slug = slug.Make(producto.Nombre)
		datos.Precio = producto.Precio
		datos.Stock = producto.Stock
		datos.Descripcion = producto.Descripcion
		datos.CategoriaID = producto.CategoriaID
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
func Productos_delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	datos := modelos.Producto{}
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
