package server

import (
	tea "github.com/charmbracelet/bubbletea"
)

type gameModel struct {
	common *commonModel
	join *joinModel
}

func NewGame(com *commonModel) *gameModel {
	g := gameModel{
		common: com,
		join: NewJoinModel(com),
	}
	return &g
}

func (m *gameModel) Init() tea.Cmd {
	return nil
}

func (m *gameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.common.choice {
	case m.common.choices[1]: // join
		j, cmd := m.join.Update(msg)
		m.join = j.(*joinModel)
		return m, cmd
	default:
		return m, nil
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