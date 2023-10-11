package server

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type joinModel struct {
	common *commonModel
	idInput textinput.Model
	tried bool
}

func NewJoinModel(com *commonModel) *joinModel {
	ii := textinput.New()
	ii.Placeholder = "XXXXXX"
	ii.CharLimit = 6
	ii.Width = 20

	j := joinModel {
		common: com,
		idInput: ii,
		tried: false,
	}

	return &j
}

func (m joinModel) Init() tea.Cmd {
    return nil
}

func (m *joinModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	cmds = append(cmds, m.idInput.Focus())

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String(){
		case "enter":
			m.idInput.Reset()
			return m, nil
		default:
			m.idInput, cmd = m.idInput.Update(msg)
			return m, cmd
		} // switch KeyMsg

	} // switch msg
	return m, nil
}

func (m *joinModel) View() string {
	s := strings.Builder{}
	s.WriteString("Enter Room Code...\n")
	s.WriteString(m.idInput.View())
	if m.tried {
		s.WriteString("\nInvalid! Try again or create your own room!")
	}
	return s.String()
}
