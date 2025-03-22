package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func getRecipeNames() []string {
	entries, err := os.ReadDir(defaultRecipeDirectory)
	if err != nil {
		os.Exit(1)
	}

	names := make([]string, len(entries))

	for i, e := range entries {
		names[i] = strings.TrimSuffix(e.Name(), ".json")
	}

	return names
}

func addRecipeIngredients(ingredients IngredientData, recipeFile string, multiplier int) {
	data, err := unmarshallRecipe(recipeFile + ".json")

	if err != nil {
		fmt.Println(err)
		panic("Problem reading/unmarshalling recipe")
	}

	for ingredient, count := range data.Ingredients {
		_, ok := ingredients[ingredient]
		if ok {
			ingredients[ingredient] += count * float32(multiplier)
		} else {
			ingredients[ingredient] = count * float32(multiplier)
		}
	}

}

func unmarshallRecipe(recipeFilename string) (*recipeData, error) {
	data, err := os.ReadFile(defaultRecipeDirectory + "/" + recipeFilename)
	if err != nil {
		return nil, err
	}

	var obj recipeData

	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func setRecipeDirectory() {
	if recipeDirectory != "" {
		return
	}

	if userRecipeDirectory != "" {
		recipeDirectory = userRecipeDirectory
		return
	}

	recipeDirectory = defaultRecipeDirectory

}

func makePageHeader(text string) string {
	var b strings.Builder

	b.WriteString(strings.Repeat("#", len(text)+4))
	b.WriteString("\n# ")
	b.WriteString(text)
	b.WriteString(" #\n")
	b.WriteString(strings.Repeat("#", len(text)+4))
	b.WriteString("\n\n")

	return b.String()
}
