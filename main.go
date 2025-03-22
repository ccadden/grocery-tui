package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var recipeDirectory = flag.String("r", defaultRecipeDirectory(), "path to directory containing recipes as JSON")

func main() {
	flag.Parse()

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
