package main

type IngredientData map[string]float32

type CartData map[string]int

type recipeData struct {
	Description string         `json:"description"`
	Ingredients IngredientData `json:"ingredients"`
}

type model struct {
	recipes     []string
	cursor      int
	cart        CartData
	ingredients IngredientData
	state       int
}
