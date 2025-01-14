package menus

import (
	"fmt"

	"github.com/abroudoux/branch/internal/ui"
	"github.com/abroudoux/branch/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func initialActionModel(branch string) actionChoice {
	actions := []string{
		"Exit",
		"Delete",
		"Merge",
		"Branch",
		"Rename",
		"Checkout",
		"Name",
	}

	return actionChoice{
		actions:        actions,
		cursor:         len(actions) - 1,
		selectedAction: "",
		selectedBranch: branch,
	}
}

type actionChoice struct {
	actions        []string
	cursor         int
	selectedAction string
	selectedBranch string
}

func (menu actionChoice) Init() tea.Cmd {
	return nil
}

func (menu actionChoice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return menu, tea.Quit
		case "down":
			menu.cursor++
			if menu.cursor >= len(menu.actions) {
				menu.cursor = 0
			}
		case "up":
			menu.cursor--
			if menu.cursor < 0 {
				menu.cursor = len(menu.actions) - 1
			}
		case "enter":
			menu.selectedAction = menu.actions[menu.cursor]
			return menu, tea.Quit
		}
	}

	return menu, nil
}

func (menu actionChoice) View() string {
	s := "\033[H\033[2J"
    s += fmt.Sprintf("Branch: %s\n\n", ui.RenderBranch(menu.selectedBranch))
	s += "Choose an action:\n\n"

	for i, action := range menu.actions {
		cursor := " "

		if menu.cursor == i {
            cursor = ui.RenderCursor()
            s += fmt.Sprintf("%s %s\n", cursor, ui.RenderActionSelected(action, true))
        } else {
            s += fmt.Sprintf("%s %s\n", cursor, ui.RenderActionSelected(action, false))
        }
	}

	return s
}

func ChooseAction(branchSelected string) (string, error) {
	actionsMenu := tea.NewProgram(initialActionModel(branchSelected))
	finalActionModel, err := actionsMenu.Run()
	if err != nil {
		return "", fmt.Errorf("error running the actions menu: %v", err)
	}

	actionMenu := finalActionModel.(actionChoice)
	return utils.CleanString(actionMenu.selectedAction), nil
}