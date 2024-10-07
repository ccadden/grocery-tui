package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	recipes []string
	cursor  int
	cart    map[string]int
	debug   bool
}

func initialModel() model {
	return model{
		recipes: []string{"a", "b", "c"},
		cart:    make(map[string]int),
	}
}

func (m model) Init() tea.Cmd {
	// thinking this is where we can read in recipe slugs from some central directory
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
		case "enter", " ":
			_, ok := m.cart[m.recipes[m.cursor]]
			if ok {
				m.cart[m.recipes[m.cursor]]++
			} else {
				m.cart[m.recipes[m.cursor]] = 1
			}
		case "p":
			m.debug = true
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.debug {
		defer func() { m.debug = false }()
		s := fmt.Sprintf("\nDump of cart:\n")

		for key, val := range m.cart {
			s += fmt.Sprintf("\n%s: %v\n", key, val)
		}

		return s
	}

	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.recipes {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// The footer
	s += "\nPress p to print.\n"
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
