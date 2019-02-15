package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/alx-t/cookbook/backend/common"
	"github.com/alx-t/cookbook/backend/data/recipe"
)

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	rs, err := recipe.Read(common.Db, "")
	j, err := json.Marshal(RecipesResource{Data: rs})

	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func GetRecipeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.DisplayAppError(w, err, "Error id converting", 500)
		return
	}

	rcp, err := recipe.ReadOne(common.Db, id)
	j, err := json.Marshal(RecipeResource{Data: rcp})

	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var dataResource RecipeResource

	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Recipe data", 500)
		return
	}

	rcp := &dataResource.Data

	if _, err := recipe.Insert(common.Db, rcp); err != nil {
		common.DisplayAppError(w, err, "Invalid Recipe data", 500)
		return
	}

	j, err := json.Marshal(RecipeResource{Data: *rcp})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		common.DisplayAppError(w, err, "Error id converting", 500)
		return
	}

	var dataResource RecipeResource
	err = json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Recipe data", 500)
		return
	}

	rcp := &dataResource.Data

	fmt.Println(id)

	if _, err := recipe.Update(common.Db, id, rcp); err != nil {
		common.DisplayAppError(w, err, "Invalid Recipe data", 500)
		return
	}

	j, err := json.Marshal(RecipeResource{Data: *rcp})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(map[string]string{"message": "Delete Recipe"})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.Write(j)
}
