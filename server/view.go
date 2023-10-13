package server

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
)

type (
	errMsg error
)

type state int

const (
	showMenu state = iota
	showGame
)

type commonModel struct {
	choices []string // items on the to-do list
	choice  string
	chosen  bool
	begin   bool
	srv     *Server
	sess    ssh.Session
	program *tea.Program
}
type parentModel struct {
	state  state
	common *commonModel
	menu   *menuModel
	game   *gameModel
}

func GetModelOption(s ssh.Session, options []string, server *Server, sess ssh.Session) *tea.Program {
	model := Model(options, server, sess)
	p := tea.NewProgram(
		model,
		tea.WithInput(s),
		tea.WithOutput(s),
	)
	model.common.program = p
	return p
}

func Model(options []string, server *Server, sess ssh.Session) *parentModel {
	common := commonModel{
		choices: options,
		choice:  "",
		chosen:  false,
		begin:   false,
		srv:     server,
		sess:    sess,
	}

	p := parentModel{
		common: &common,
		menu:   NewMenu(&common),
		game:   NewGame(&common),
	}

	return &p
}

func (m *parentModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m *parentModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		var cmd tea.Cmd
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "ctrl+n":
			return m, cmd
		}
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		slog.Info("change size", "width", msg.Width, "height", msg.Height)
	}

	switch m.state {
	case showMenu:
		var cmd2 tea.Cmd
		men, cmd := m.menu.Update(msg)
		m.menu = men.(*menuModel)
		if m.common.chosen {
			m.state = showGame
			var g tea.Model
			g, cmd2 = m.game.Update(msg)
			m.game = g.(*gameModel)
		}
		return m, tea.Batch(cmd, cmd2)
	case showGame:
		g, cmd := m.game.Update(msg)
		m.game = g.(*gameModel)
		return m, cmd
	}
	return m, nil
}

func (m *parentModel) View() string {
	switch m.state {
	case showMenu:
		return lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(21)).Render(m.menu.View())
	case showGame:
		return lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(21)).Render(m.game.View())
	}
	return ""
}
