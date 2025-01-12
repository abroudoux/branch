package menus

import (
	"fmt"

	"github.com/abroudoux/branch/internal/ui"
	"github.com/abroudoux/branch/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type BranchChoice struct {
	branches        []string
	cursor          int
	selectedBranch  string
}

func initialBranchModel(branches []string) BranchChoice {
	return BranchChoice{
		branches:        branches,
		cursor:          len(branches) - 1,
		selectedBranch:  "",
	}
}

func (menu BranchChoice) Init() tea.Cmd {
	return nil
}

func (menu BranchChoice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			menu.selectedBranch = menu.branches[menu.cursor]
			return menu, tea.Quit
		}
	}

	return menu, nil
}

func (menu BranchChoice) View() string {
    s := "\033[H\033[2J"
    s += "Choose a branch:\n\n"

    for i, branch := range menu.branches {
        cursor := " "

        if menu.cursor == i {
            cursor = ui.RenderCursor()
            s += fmt.Sprintf("%s %s\n", cursor, ui.RenderBranchSelected(branch, true))
        } else {
            s += fmt.Sprintf("%s %s\n", cursor, ui.RenderBranchSelected(branch, false))
        }
    }

    return s
}

func ChooseBranch(branches []string) (string, error) {
	branchesMenu := tea.NewProgram(initialBranchModel(branches))
	finalModel, err := branchesMenu.Run()
	if err != nil {
		return "", fmt.Errorf("error running the branches menu: %v", err)
	}

	branchMenu := finalModel.(BranchChoice)
	return utils.CleanString(branchMenu.selectedBranch), nil
}