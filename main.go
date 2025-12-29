package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ---------------------------------------------------------
// 1. MODEL
// The model stores the application state.
// ---------------------------------------------------------

type model struct {
	items    []string
	selected map[int]struct{}
	cursor   int     
}

func initialModel() model {
	return model{
		items:    []string{"Read r/golang", "Post on r/commandline", "Upvote cool TUIs", "Check Karma"},
		selected: make(map[int]struct{}),
		cursor:   0,
	}
}

// ---------------------------------------------------------
// 2. INIT
// Init performs initial I/O (like a timer). Usually nil for simple apps.
// ---------------------------------------------------------

func (m model) Init() tea.Cmd {
	return nil
}

// ---------------------------------------------------------
// 3. UPDATE
// Update handles messages (keypresses, window resizes, etc.)
// ---------------------------------------------------------

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
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model and no command
	return m, nil
}

// ---------------------------------------------------------
// 4. VIEW
// View renders the UI string based on the current model state.
// ---------------------------------------------------------

func (m model) View() string {
	s := "Reddit TUI Checklist\n\n"
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	checkedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Strikethrough(true)

	for i, item := range m.items {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = cursorStyle.Render(">") // cursor present
		}

		checked := " " // not checked
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		line := fmt.Sprintf("%s [%s] %s", cursor, checked, item)

		if _, ok := m.selected[i]; ok {
			line = fmt.Sprintf("%s [%s] %s", cursor, checked, checkedStyle.Render(item))
		}

		s += line + "\n"
	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error occured %v", err)
		os.Exit(1)
	}
}