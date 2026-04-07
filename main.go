package main

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1).
			MarginBottom(1)

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true)

	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))

	dimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888"))

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA")).
			Italic(true)
)

type model struct {
	configs  []Config
	cursor   int
	selected bool
	chosen   Config
}

type quitMsg struct {
	config Config
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
			if m.cursor < len(m.configs)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = true
			m.chosen = m.configs[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("SSH Launcher") + "\n\n"

	for i, cfg := range m.configs {
		cursor := " "
		style := normalStyle
		dim := dimStyle

		if m.cursor == i {
			cursor = ">"
			style = selectedStyle
			dim = normalStyle
		}

		userDisplay := cfg.User
		if userDisplay == "" {
			userDisplay = "default"
		}

		// Format: > Title (user@host)
		s += fmt.Sprintf("%s %s  %s(%s@%s)\n",
			cursor,
			style.Render(cfg.Title),
			dim.Render("Connection:"),
			dim.Render(userDisplay),
			dim.Render(cfg.Host),
		)
	}

	s += "\n(q to quit, enter to connect)\n"
	return s
}

func main() {
	configs, err := loadConfigs()
	if err != nil {
		if err.Error() == "CONFIG_NOT_FOUND" {
			fmt.Println("Error: Configuration file not found.")
			fmt.Printf("Please create a file at: %s\n", getConfigPath())
			fmt.Println("\nExample format:")
			fmt.Println(getExampleConfig())
		} else {
			fmt.Printf("Error loading configs: %v\n", err)
		}
		os.Exit(1)
	}

	if len(configs) == 0 {
		fmt.Println("No hosts configured in ssh_launcher.json")
		os.Exit(1)
	}

	p := tea.NewProgram(model{
		configs: configs,
		cursor:  0,
	})

	finalModel, err := p.Run()
	if err != nil {
		fmt.Printf("Error running TUI: %v\n", err)
		os.Exit(1)
	}

	m := finalModel.(model)

	if m.selected {
		launchSSH(m.chosen)
	}
}

func launchSSH(cfg Config) {
	var args []string
	if cfg.User != "" {
		args = []string{fmt.Sprintf("%s@%s", cfg.User, cfg.Host)}
	} else {
		args = []string{cfg.Host}
	}

	cmd := exec.Command("ssh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing ssh: %v\n", err)
		os.Exit(1)
	}
}
