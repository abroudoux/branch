package branches

import (
	"fmt"

	"github.com/abroudoux/branch/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func initialBranchChoiceModel(branches []BranchDetails) branchChoice {
	return branchChoice{
		branches:       branches,
		cursor:         len(branches) - 1,
		branchSelected: BranchDetails{},
	}
}

func (menu branchChoice) Init() tea.Cmd {
	return nil
}

func (menu branchChoice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return menu, tea.Quit
		case "down":
			menu.cursor++
			if menu.cursor >= len(menu.branches) {
				menu.cursor = 0
			}
		case "up":
			menu.cursor--
			if menu.cursor < 0 {
				menu.cursor = len(menu.branches) - 1
			}
		case "enter":
			menu.branchSelected = menu.branches[menu.cursor]
			return menu, tea.Quit
		}
	}

	return menu, nil
}

func (menu branchChoice) View() string {
	s := "\033[H\033[2J\n"
	s += "Choose a branch:\n\n"

	for i, branch := range menu.branches {
		cursor := ui.RenderCursor(menu.cursor == i)
		if branch.IsHead {
			branchName := "* " + branch.Name
			s += fmt.Sprintf("%s %s\n", cursor, ui.RenderCurrentLine(branchName, menu.cursor == i))
		} else {
			branchName := "  " + branch.Name
			s += fmt.Sprintf("%s %s\n", cursor, ui.RenderCurrentLine(branchName, menu.cursor == i))
		}
	}

	s += "\n"

	return s
}

func SelectBranch(branches []BranchDetails) (BranchDetails, error) {
	p := tea.NewProgram(initialBranchChoiceModel(branches))
	m, err := p.Run()
	if err != nil {
		return BranchDetails{}, err
	}

	branchSelected := m.(branchChoice).branchSelected
	return branchSelected, nil
}
