package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func initialModel() model {
	return model{
		recipes:     getRecipeNames(),
		cart:        make(CartData),
		state:       Shopping,
		ingredients: make(IngredientData),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
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
		s = fmt.Sprintf("\nThis Week's Cart:\n")

		for key, val := range m.cart {
			s += fmt.Sprintf("\n%s: %v\n", key, val)
		}

		s += m.buildInstructionFooter()

		return s
	case Shopping:
		s += makePageHeader("What's for Dinner?")

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
		s += m.buildInstructionFooter()

		// Send the UI for rendering
		return s
	case Checkout:
		s += makePageHeader("Your Cart")

		for key, val := range m.cart {
			s += fmt.Sprintf("%s %v\n", key, val)
		}

		s += m.buildInstructionFooter()

		return s
	case List:
		s += makePageHeader("To Buy This Week")

		for key, val := range m.ingredients {
			s += fmt.Sprintf("%s %v\n", key, val)
		}

		s += m.buildInstructionFooter()

		return s

	default:
		return s
	}
}

func (m *model) buildInstructionFooter() string {
	switch m.state {
	case Shopping:
		return fmt.Sprintf("\n%s\n\n%s\n\n%s\n", shoppingInstructionsString, checkoutString, quitString)
	case ViewCart:
		return fmt.Sprintf("\n%s\n%s\n\n%s\n", returnToShoppingString, checkoutString, quitString)
	case Checkout:
		return fmt.Sprintf("\n%s\n%s\n\n%s\n", shoppingListString, returnToShoppingString, quitString)
	case List:
		return fmt.Sprintf("\n%s\n\n%s\n", returnToShoppingString, quitString)
	}

	return ""
}
