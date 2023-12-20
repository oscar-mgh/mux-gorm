package main

import (
	"clase_5_mux_gorm/modelos"
	"clase_5_mux_gorm/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	modelos.Migraciones()
	r := mux.NewRouter()
	routes.UseRoutes(r)
	handler := cors.AllowAll().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
