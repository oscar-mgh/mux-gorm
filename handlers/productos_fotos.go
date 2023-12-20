package handlers

import (
	"clase_5_mux_gorm/database"
	"clase_5_mux_gorm/modelos"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func ProductosFotos_get_por_producto(w http.ResponseWriter, r *http.Request) {
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
		fotos := modelos.ProductoFotos{}
		database.Database.Where("producto_id=?", id).Find(&fotos)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(fotos)
	}
}
func ProductosFotosUpload(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("foto")
	var extension = strings.Split(handler.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	foto := string(time[4][6:14]) + "." + extension
	var archivo string = "public/uploads/productos/" + foto
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}
	//crear el registro en la base de datos
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	datos := modelos.ProductoFoto{Nombre: foto, ProductoID: id}
	database.Database.Save(&datos)
	w.Header().Set("Content-Type", "application/json")
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Se creó el registro exitosamente ",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
func ProductosFotosDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	datos := modelos.ProductoFoto{}
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
		e := os.Remove("public/uploads/productos/" + datos.Nombre)
		if e != nil {
			log.Fatal(e)
		}
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
