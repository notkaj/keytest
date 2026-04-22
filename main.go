package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	table "charm.land/lipgloss/v2/table"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	return model{}
}

type model struct {
	text      string
	keystroke string
	string    string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok {
		return model{text: key.Text, keystroke: key.Keystroke(), string: key.String()}, nil
	}
	return m, nil
}

func (m model) View() tea.View {
	rows := [][]string{
		{"text", m.text},
		{"keystroke", m.keystroke},
		{"string", m.string},
	}
	table := table.New().Border(lipgloss.RoundedBorder()).Rows(rows...).Headers("field", "result")
	return tea.NewView(table.Render())
}
