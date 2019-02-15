package controllers

import (
	"github.com/alx-t/cookbook/backend/models"
)

type (
	// For Get - /recipes
	RecipesResource struct {
		Data []models.Recipe `json:"recipes"`
	}

	// For Post/Put - /recipes
	// For Get - /recipes/id
	RecipeResource struct {
		Data models.Recipe `json:"recipe"`
	}
)
