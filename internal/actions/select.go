package actions

import (
	"fmt"

	"github.com/abroudoux/branch/internal/branches"
	"github.com/abroudoux/branch/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func (a BranchAction) String() string {
	return [...]string{
		"Exit",
		"Delete",
		"Merge",
		"New Branch",
		"Checkout",
		"Copy Name",
	}[a]
}

func getAllBranchActions() []BranchAction {
	return []BranchAction{
		BranchActionExit,
		BranchActionDelete,
		BranchActionMerge,
		BranchActionNewBranch,
		BranchActionCheckout,
		BranchActionCopyName,
	}
}

func initialBranchActionChoiceModel(branch branches.BranchDetails) branchActionChoice {
	allBranchesActions := getAllBranchActions()

	return branchActionChoice{
		actions:        allBranchesActions,
		cursor:         len(allBranchesActions) - 1,
		actionSelected: BranchActionExit,
		branchSelected: branch,
	}
}

func (menu branchActionChoice) Init() tea.Cmd {
	return nil
}

func (menu branchActionChoice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			menu.actionSelected = menu.actions[menu.cursor]
			return menu, tea.Quit
		}
	}

	return menu, nil
}

func (menu branchActionChoice) View() string {
	s := "\033[H\033[2J\n"
	s += fmt.Sprintf("Choose an action for the branch %s:\n\n", ui.RenderElementSelected(string(menu.branchSelected.Name)))

	for i, action := range menu.actions {
		cursor := ui.RenderCursor(menu.cursor == i)
		s += fmt.Sprintf("%s %s\n", cursor, ui.RenderCurrentLine(action.String(), menu.cursor == i))
	}

	s += "\n"

	return s
}

func SelectAction(branchSelected branches.BranchDetails) (BranchAction, error) {
	p := tea.NewProgram(initialBranchActionChoiceModel(branchSelected))
	m, err := p.Run()
	if err != nil {
		return BranchActionExit, err
	}

	actionSelected := m.(branchActionChoice).actionSelected
	return actionSelected, nil
}
