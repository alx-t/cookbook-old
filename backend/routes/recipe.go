package routers

import (
	"github.com/alx-t/cookbook/backend/controllers"
	"github.com/gorilla/mux"
)

func SetRecipeRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.Root).Methods("GET")

	router.HandleFunc("/recipes", controllers.NotImplemented).Methods("GET")
	router.HandleFunc("/recipes", controllers.NotImplemented).Methods("POST")
	router.HandleFunc("/recipes/{id}", controllers.NotImplemented).Methods("GET")
	router.HandleFunc("/recipes/{id}", controllers.NotImplemented).Methods("PATCH")
	router.HandleFunc("/recipes/{id}", controllers.NotImplemented).Methods("DELETE")

	return router
}
