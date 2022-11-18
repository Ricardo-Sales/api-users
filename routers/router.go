package routers

import (
	"net/http"

	"github.com/Ricardo-Sales/api-users/controllers"
	"github.com/gorilla/mux"
)

func Generate() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", controllers.GetUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios", controllers.PostUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios/{id}", controllers.PutUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", controllers.GetUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", controllers.DeleteUsuario).Methods(http.MethodDelete)

	return router

}
