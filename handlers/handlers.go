package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type ResponseGenerico struct {
	Estado  string `json:"estado"`
	Mensaje string `json:"mensaje"`
}

func Ejemplo_upload(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("foto")
	var extension = strings.Split(handler.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	foto := string(time[4][6:14]) + "." + extension
	var archivo string = "public/uploads/fotos/" + foto
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
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "Se creó el archivo exitosamente",
		"foto":    foto,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
func EjemploVerFoto(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if len(file) < 1 {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	OpenFile, err := os.Open("public/uploads/" + r.URL.Query().Get("folder") + "/" + file)
	if err != nil {
		respuesta := map[string]string{
			"estado":  "error",
			"mensaje": "Ocurrió un error inesperado",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(respuesta)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar el archivo", http.StatusBadRequest)
	}
}
