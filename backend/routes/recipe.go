package routers

import (
	"github.com/alx-t/cookbook/backend/controllers"
	"github.com/gorilla/mux"
)

func SetRecipeRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.Root).Methods("GET")

	router.HandleFunc("/recipes", controllers.GetRecipes).Methods("GET")
	router.HandleFunc("/recipes", controllers.CreateRecipe).Methods("POST")
	router.HandleFunc("/recipes/{id}", controllers.GetRecipeById).Methods("GET")
	router.HandleFunc("/recipes/{id}", controllers.UpdateRecipe).Methods("PATCH")
	router.HandleFunc("/recipes/{id}", controllers.DeleteRecipe).Methods("DELETE")

	return router
}
