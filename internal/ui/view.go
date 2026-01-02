package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func renderPane(content string, width, height int, borderColor string, active bool) string {
	innerWidth := width - 2
	innerHeight := height - 2

	if innerWidth < 1 {
		innerWidth = 1
	}
	if innerHeight < 1 {
		innerHeight = 1
	}

	lines := strings.Split(content, "\n")

	result := make([]string, innerHeight)
	for i := 0; i < innerHeight; i++ {
		if i < len(lines) {
			line := lines[i]

			if lipgloss.Width(line) > innerWidth {
				runes := []rune(line)
				if len(runes) > innerWidth {
					line = string(runes[:innerWidth])
				}
			}

			result[i] = line + strings.Repeat(" ", innerWidth-lipgloss.Width(line))
		} else {
			result[i] = strings.Repeat(" ", innerWidth)
		}
	}

	innerContent := strings.Join(result, "\n")

	color := lipgloss.Color(borderColor)
	if active {
		color = lipgloss.Color("205")
	}

	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(color).
		Width(innerWidth).
		Height(innerHeight)

	return style.Render(innerContent)
}

func (m Model) View() string {
	if m.Width == 0 || m.Height == 0 {
		return ""
	}

	controlPaneHeight := 3

	sidebarWidth := m.Width / 5
	if sidebarWidth < 15 {
		sidebarWidth = 15
	}
	remainingWidth := m.Width - sidebarWidth
	postsWidth := remainingWidth / 2
	previewWidth := remainingWidth - postsWidth

	paneHeight := m.Height - controlPaneHeight

	postsPaneHeading := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true).MarginLeft(2)
	navPaneHeading := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true).MarginLeft(2)
	previewPaneHeading := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true).MarginLeft(2)
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	postTitleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("86"))
	subredditStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("33"))
	metaStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

	sidebarContent := navPaneHeading.Render("NAVIGATION") + "\n\n"
	for i, item := range m.SidebarItems {
		cursor := "  "
		if m.SidebarCursor == i {
			cursor = cursorStyle.Render("> ")
		}
		sidebarContent += cursor + item + "\n"
	}

	postsContent := postsPaneHeading.Render("POSTS") + "\n\n"
	for i, post := range m.Posts {
		cursor := "  "
		if m.PostsCursor == i {
			cursor = cursorStyle.Render("> ")
		}
		postsContent += cursor + postTitleStyle.Render(post.Title) + "\n"
		postsContent += "   " + subredditStyle.Render(post.Subreddit) + " by u/" + post.Author + "\n"
		postsContent += "   " + metaStyle.Render(fmt.Sprintf("%d upvotes | %d comments", post.Upvotes, post.Comments)) + "\n\n"
	}

	var previewContent string
	if m.PostsCursor >= 0 && m.PostsCursor < len(m.Posts) {
		selectedPost := m.Posts[m.PostsCursor]
		previewContent = previewPaneHeading.Render("PREVIEW") + "\n\n"
		previewContent += postTitleStyle.Render(selectedPost.Title) + "\n\n"
		previewContent += subredditStyle.Render(selectedPost.Subreddit) + " by u/" + selectedPost.Author + "\n"
		previewContent += metaStyle.Render(fmt.Sprintf("%d upvotes | %d comments", selectedPost.Upvotes, selectedPost.Comments)) + "\n\n"
		previewContent += strings.Repeat("-", 20) + "\n\n"
		previewContent += "Lorem ipsum dolor sit amet,\n"
		previewContent += "consectetur adipiscing elit.\n\n"
		previewContent += "Sed do eiusmod tempor incididunt\n"
		previewContent += "ut labore et dolore magna aliqua."
	} else {
		previewContent = "PREVIEW\n\nSelect a post to view"
	}

	sidebar := renderPane(sidebarContent, sidebarWidth, paneHeight, "63", m.ActivePane == "sidebar")
	posts := renderPane(postsContent, postsWidth, paneHeight, "63", m.ActivePane == "posts")
	preview := renderPane(previewContent, previewWidth, paneHeight, "63", false)

	mainContent := lipgloss.JoinHorizontal(lipgloss.Top, sidebar, posts, preview)

	controlText := metaStyle.Render("Tab: switch panes | ↑↓/j/k: navigate | q: quit")
	controlPane := renderPane(controlText, m.Width, controlPaneHeight, "63", false)

	return lipgloss.JoinVertical(lipgloss.Left, mainContent, controlPane)
}
