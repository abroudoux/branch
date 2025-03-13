package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func getConfirmation(msg string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (y/n) [yes]: ", msg)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	confirmation := strings.TrimSpace(input)
	if confirmation == "" || strings.EqualFold(confirmation, "y") || strings.EqualFold(confirmation, "yes") {
		return true, nil
	}

	return false, nil
}

func readInput(message string) (string, error) {
	var input string

	fmt.Print(message)

	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}

	return input, nil
}

func initialBranchChoiceModel(branches []Branch, head Branch) branchChoice {
	return branchChoice{
		head:           head,
		branches:       branches,
		cursor:         len(branches) - 1,
		branchSelected: nil,
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
			menu.branchSelected = nil
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
		cursor := renderCursor(menu.cursor == i)

		if branch.Name() == menu.head.Name() {
			branchName := "* " + branch.Name().Short()
			s += fmt.Sprintf("%s %s\n", cursor, renderCurrentLine(branchName, menu.cursor == i))
		} else {
			branchName := "  " + branch.Name().Short()
			s += fmt.Sprintf("%s %s\n", cursor, renderCurrentLine(branchName, menu.cursor == i))
		}
	}

	s += "\n"

	return s
}

const (
	BranchActionExit branchAction = iota
	BranchActionDelete
	BranchActionMerge
	BranchActionNewBranch
	BranchActionCheckout
	BranchActionCopyName
)

func (a branchAction) String() string {
	return [...]string{
		"Exit",
		"Delete",
		"Merge",
		"New Branch",
		"Checkout",
		"Copy Name",
	}[a]
}

func getAllBranchActions() []branchAction {
	return []branchAction{
		BranchActionExit,
		BranchActionDelete,
		BranchActionMerge,
		BranchActionNewBranch,
		BranchActionCheckout,
		BranchActionCopyName,
	}
}

func initialBranchActionChoiceModel(branch Branch) branchActionChoice {
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
	s += fmt.Sprintf("Choose an action for the branch %s:\n\n", renderElSelected(string(menu.branchSelected.Name().Short())))

	for i, action := range menu.actions {
		cursor := renderCursor(menu.cursor == i)
		s += fmt.Sprintf("%s %s\n", cursor, renderCurrentLine(action.String(), menu.cursor == i))
	}

	s += "\n"

	return s
}

func SelectAction(branchSelected Branch) (branchAction, error) {
	p := tea.NewProgram(initialBranchActionChoiceModel(branchSelected))
	m, err := p.Run()
	if err != nil {
		return BranchActionExit, err
	}

	actionSelected := m.(branchActionChoice).actionSelected
	return actionSelected, nil
}
