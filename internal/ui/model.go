package ui

import (
	"reddit-tui/internal/models"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	SidebarItems  []string
	Posts         []models.Post
	SidebarCursor int
	PostsCursor   int
	ActivePane    string
	Width         int
	Height        int
}

func InitialModel() Model {
	return Model{
		SidebarItems: []string{"Home", "Popular", "Explore", "Settings", "Login/Auth"},
		Posts: []models.Post{
			{Title: "Building a Reddit TUI with Go and Bubble Tea", Subreddit: "r/golang", Author: "gopher_dev", Upvotes: 342, Comments: 45},
			{Title: "What are your favorite terminal tools?", Subreddit: "r/commandline", Author: "cli_enthusiast", Upvotes: 528, Comments: 89},
			{Title: "Show HN: My weekend project - a Reddit client for the terminal", Subreddit: "r/programming", Author: "weekend_coder", Upvotes: 1205, Comments: 134},
			{Title: "TUI vs GUI: The eternal debate", Subreddit: "r/linux", Author: "terminal_lover", Upvotes: 876, Comments: 201},
			{Title: "Charm libraries are amazing for building TUIs", Subreddit: "r/golang", Author: "bubble_fan", Upvotes: 445, Comments: 67},
			{Title: "Ask Reddit: What's your development setup?", Subreddit: "r/AskReddit", Author: "curious_dev", Upvotes: 2301, Comments: 456},
			{Title: "Vim vs Emacs: A comprehensive comparison", Subreddit: "r/programming", Author: "editor_wars", Upvotes: 689, Comments: 342},
			{Title: "Why I switched from GUI apps to terminal", Subreddit: "r/commandline", Author: "minimalist_dev", Upvotes: 934, Comments: 178},
		},
		SidebarCursor: 0,
		PostsCursor:   0,
		ActivePane:    "sidebar",
		Width:         80,
		Height:        24,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.ActivePane == "sidebar" {
				m.ActivePane = "posts"
			} else {
				m.ActivePane = "sidebar"
			}
		case "up", "k":
			if m.ActivePane == "sidebar" {
				if m.SidebarCursor > 0 {
					m.SidebarCursor--
				}
			} else {
				if m.PostsCursor > 0 {
					m.PostsCursor--
				}
			}
		case "down", "j":
			if m.ActivePane == "sidebar" {
				if m.SidebarCursor < len(m.SidebarItems)-1 {
					m.SidebarCursor++
				}
			} else {
				if m.PostsCursor < len(m.Posts)-1 {
					m.PostsCursor++
				}
			}
		}
	}

	return m, nil
}
