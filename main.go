package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var recipeDirectory string = "./recipes"

const (
	Shopping = iota
	ViewCart
	Checkout
	List
)

type model struct {
	recipes     []string
	cursor      int
	cart        map[string]int
	ingredients map[string]int
	state       int
}

type IngredientData map[string]int

type recipeData struct {
	Description string         `json:"description"`
	Ingredients IngredientData `json:"ingredients"`
}

func getRecipeNames() []string {
	entries, err := os.ReadDir(recipeDirectory)
	if err != nil {
		os.Exit(1)
	}

	names := make([]string, len(entries))

	for i, e := range entries {
		names[i] = strings.TrimSuffix(e.Name(), ".json")
	}

	return names
}

func initialModel() model {
	return model{
		recipes:     getRecipeNames(),
		cart:        make(map[string]int),
		state:       Shopping,
		ingredients: make(map[string]int),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.recipes)-1 {
				m.cursor++
			}
		case " ", "a":
			switch m.state {
			case Shopping:
				_, ok := m.cart[m.recipes[m.cursor]]
				if ok {
					m.cart[m.recipes[m.cursor]]++
				} else {
					m.cart[m.recipes[m.cursor]] = 1
				}
			}
		case "enter":
			switch m.state {
			case Shopping:
				_, ok := m.cart[m.recipes[m.cursor]]
				if ok {
					m.cart[m.recipes[m.cursor]]++
				} else {
					m.cart[m.recipes[m.cursor]] = 1
				}
			case Checkout:
				for recipe, count := range m.cart {
					addRecipeIngredients(m.ingredients, recipe, count)
				}

				m.state = List
			}

		case "backspace", "delete", "x":
			_, ok := m.cart[m.recipes[m.cursor]]
			if ok {
				m.cart[m.recipes[m.cursor]]--
			} else {
				m.cart[m.recipes[m.cursor]] = 0
			}
		case "p":
			m.state = ViewCart
		case "s":
			m.state = Shopping
		case "g":
			m.state = Checkout
		}
	}

	return m, nil
}

func (m model) View() string {
	s := ""
	switch m.state {
	case ViewCart:
		s = fmt.Sprintf("\nDump of cart:\n")

		for key, val := range m.cart {
			s += fmt.Sprintf("\n%s: %v\n", key, val)
		}

		s += "\nPress s to return to browsing.\n"

		return s
	case Shopping:
		// The header
		s += "What are we eating this week?\n\n"

		// Iterate over our choices
		for i, choice := range m.recipes {

			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
			}

			// Render the row
			s += fmt.Sprintf("%s [%v] %s\n", cursor, m.cart[choice], choice)
		}

		// The footer
		s += "\nPress p to print.\n"
		s += "\nPress g to checkout.\n"
		s += "\nPress q to quit.\n"

		// Send the UI for rendering
		return s
	case Checkout:
		s += "Your Cart\n\n"

		for key, val := range m.cart {
			s += fmt.Sprintf("%s %v\n", key, val)
		}

		s += "Ingredients debugging\n\n"
		for key, val := range m.ingredients {
			s += fmt.Sprintf("%s, %v\n", key, val)
		}

		// The footer
		s += "\nPress Enter to create shopping list.\n"
		s += "\nPress s to return to shopping.\n"
		s += "\nPress q to quit.\n"

		// Send the UI for rendering
		return s
	case List:
		s += "To Buy This Week\n\n"

		// Iterate over our choices
		for key, val := range m.ingredients {
			s += fmt.Sprintf("%s %v\n", key, val)
		}

		// The footer
		s += "\nPress q to quit.\n"
		return s

	default:
		return s
	}
}

func addRecipeIngredients(ingredients map[string]int, recipeFile string, multiplier int) {
	data, err := unmarshallRecipe(recipeFile + ".json")

	if err != nil {
		fmt.Println(err)
		panic("Problem reading/unmarshalling recipe")
	}

	for ingredient, count := range data.Ingredients {
		_, ok := ingredients[ingredient]
		if ok {
			ingredients[ingredient] += count * multiplier
		} else {
			ingredients[ingredient] = count * multiplier
		}
	}

}

func unmarshallRecipe(recipeFilename string) (*recipeData, error) {
	data, err := os.ReadFile(recipeDirectory + "/" + recipeFilename)
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

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
