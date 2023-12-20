package routes

import (
	"clase_5_mux_gorm/handlers"
	"clase_5_mux_gorm/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func UseRoutes(mux *mux.Router) http.Handler {
	mux.PathPrefix("/api")
	mux.HandleFunc("upload", handlers.Ejemplo_upload).Methods("POST")
	mux.HandleFunc("archivo", handlers.EjemploVerFoto).Methods("GET")

	mux.HandleFunc("categorias", handlers.Categoria_get).Methods("GET")
	mux.HandleFunc("categorias/{id:[0-9]+}", handlers.Categoria_get_con_parametro).Methods("GET")
	mux.HandleFunc("categorias", handlers.Categoria_post).Methods("POST")
	mux.HandleFunc("categorias/{id:[0-9]+}", handlers.Categoria_put).Methods("PUT")
	mux.HandleFunc("categorias/{id:[0-9]+}", handlers.Categoria_delete).Methods("DELETE")

	mux.HandleFunc("productos", handlers.Productos_get).Methods("GET")
	mux.HandleFunc("productos/{id:[0-9]+}", handlers.Productos_get_con_parametro).Methods("GET")
	mux.HandleFunc("productos", handlers.Productos_post).Methods("POST")
	mux.HandleFunc("productos/{id:[0-9]+}", handlers.Productos_put).Methods("PUT")
	mux.HandleFunc("productos/{id:[0-9]+}", handlers.Productos_delete).Methods("DELETE")

	mux.HandleFunc("productos-fotos/{id:[0-9]+}", handlers.ProductosFotosUpload).Methods("POST")
	mux.HandleFunc("productos-fotos/{id:[0-9]+}", handlers.ProductosFotos_get_por_producto).Methods("GET")
	mux.HandleFunc("productos-fotos/{id:[0-9]+}", handlers.ProductosFotosDelete).Methods("DELETE")

	mux.HandleFunc("seguridad/registro", handlers.Seguridad_registro).Methods("POST")
	mux.HandleFunc("seguridad/login", handlers.Seguridad_login).Methods("POST")
	mux.HandleFunc("seguridad/protegido", middleware.ValidarJWT(handlers.Seguridad_protegido)).Methods("GET")

	return mux
}
