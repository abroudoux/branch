package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"

	git "github.com/abroudoux/branch/internal/git"
	repository "github.com/abroudoux/branch/internal/repository"
	ui "github.com/abroudoux/branch/internal/ui"
	"github.com/abroudoux/branch/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Args) > 1 {
		err := flagMode()
		if err != nil {
			utils.PrintErrorAndExit(err)
		}
		os.Exit(0)
	}

	err := isInGitRepository()
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	branches := git.GetBranchesWithDefaultIndication()

	branch, err := chooseBranch(branches)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	action, err := chooseAction(branch)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	err = doAction(branch, action)
	if err != nil {
		utils.PrintErrorAndExit(err)
	}
}

func isGitInstalled() error {
	cmd := exec.Command("git", "version")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func isInGitRepository() error {
	err := isGitInstalled()
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func printBranches() {
	branches := git.GetBranchesWithDefaultIndication()

	for _, branch := range branches {
		fmt.Println(branch)
	}
}

func cleanString(s string) string {
	return strings.TrimSpace(strings.TrimPrefix(s, "*"))
}

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

type actionChoice struct {
	actions        []string
	cursor         int
	selectedAction string
	selectedBranch string
}

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

func chooseBranch(branches []string) (string, error) {
	branchesMenu := tea.NewProgram(initialBranchModel(branches))
	finalModel, err := branchesMenu.Run()
	if err != nil {
		return "", fmt.Errorf("error running the branches menu: %v", err)
	}

	branchMenu := finalModel.(BranchChoice)
	return cleanString(branchMenu.selectedBranch), nil
}

func chooseAction(selectedBranch string) (string, error) {
	actionsMenu := tea.NewProgram(initialActionModel(selectedBranch))
	finalActionModel, err := actionsMenu.Run()
	if err != nil {
		return "", fmt.Errorf("error running the actions menu: %v", err)
	}

	actionMenu := finalActionModel.(actionChoice)
	return cleanString(actionMenu.selectedAction), nil
}

func flagMode() error {
	flag := os.Args[1]

	switch flag {
	case "run", "-r":
		branches:= git.GetBranchesWithDefaultIndication()
		chooseBranch(branches)
	case "-v", "--version":
		latestVersion, err := repository.GetLatestRelease()
		if err != nil {
			return fmt.Errorf("error getting latest version: %v", err)
		}

		fmt.Printf("Latest version: %s\n", latestVersion)
	case "-l", "--list":
		printBranches()
	case "-h", "--help":
		repository.PrintHelpManual()
	default:
		repository.PrintHelpManual()
	}

	return nil
}

func doAction(branch string, action string) error {
	switch action {
	case "Exit":
		fmt.Println("Exiting...")
		return nil
	case "Delete":
		return git.DeleteBranch(branch)
	case "Merge":
		return git.MergeBranch(branch)
	case "Branch":
		return git.CreateBranch(branch)
	case "Rename":
		return git.RenameBranch(branch)
	case "Checkout":
		return git.CheckoutBranch(branch)
	case "Name":
		return git.CopyName(branch)
	default:
		return fmt.Errorf("invalid action: %s", action)
	}
}