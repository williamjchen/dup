package server

import (
	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
)

type gameModel struct {
	common *commonModel
	join   *joinModel
}

func NewGame(com *commonModel) *gameModel {
	g := gameModel{
		common: com,
		join:   NewJoinModel(com),
	}
	return &g
}

func (m *gameModel) Init() tea.Cmd {
	return nil
}

func (m *gameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.common.choice {
	case m.common.choices[1]: // join
		//m.join = NewJoinModel(m.common)
		if !m.join.idInput.Focused() {
			cmd = tea.Batch(m.join.idInput.Cursor.SetMode(cursor.CursorBlink),
				m.join.idInput.Focus(),
			)
		}

		j, cmd2 := m.join.Update(msg)
		m.join = j.(*joinModel)
		return m, tea.Batch(cmd, cmd2)
	default:
		return m, cmd
	}
}

func (m gameModel) View() string {
	switch m.common.choice {
	case m.common.choices[1]: // join
		return m.join.View()
	default:
		return ""
	}
}
