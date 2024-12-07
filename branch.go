package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]

		if arg == "run" || arg == "-r" {
			chooseBranch()
		} else  if arg == "-v" || arg == "--verbose" {
			fmt.Println("2.0.0")
		} else if arg == "-l" || arg == "--list" {
			printBranches()
		} else if arg == "-h" || arg == "--help" {
			printHelpManual()
		}
	}

	isGitInstalled()
	isInGitRepository()

	interactiveMenu := tea.NewProgram(initialModel())

	finalModel, err := interactiveMenu.Run()

	if err != nil {
		fmt.Printf("Error running the interactive menu: %v\n", err)
		os.Exit(1)
	}

	branchMenu := finalModel.(branchChoice)

	fmt.Printf("You selected: %s\n", branchMenu.selectedBranch)
}

func isGitInstalled() bool {
	cmd := exec.Command("git", "version")
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error checking if git is installed", err)
		os.Exit(1)
	}

	return true
}

func isInGitRepository() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error checking if in git repository", err)
		os.Exit(1)
	}

	return true
}

func getBranches() []string {
	cmd := exec.Command("git", "branch", "--format=%(refname:short)")
	branches, err := cmd.Output()

	if err != nil {
		fmt.Println("Error getting branches", err)
		os.Exit(1)
	}

	return strings.Fields(string(branches))
}

func getDefaultBranch() string {
	cmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")
	defaultBranch, err := cmd.Output()

	if err != nil {
		fmt.Println("Error getting default branch", err)
		os.Exit(1)
	}

	return strings.TrimSpace(string(defaultBranch))
}

func getBranchesWithDefaultIndication() []string {
	branches := getBranches()
	defaultBranch := getDefaultBranch()
	branchesWithDefaultIndication := []string{}

	for _, branch := range branches {
		if branch == defaultBranch {
			branchesWithDefaultIndication = append(branchesWithDefaultIndication, "* " + branch)
		} else {
			branchesWithDefaultIndication = append(branchesWithDefaultIndication, "  " + branch)
		}
	}

	return branchesWithDefaultIndication
}

func printHelpManual() {
	fmt.Println("Usage: branch [options]")
	fmt.Println("Options:")
	fmt.Println("branch [run | -r]        Start the interactive branch selection")
	fmt.Println("branch [--list | -l]     List all branches")
	fmt.Println("branch [--help | -h]     Show this help message")

	os.Exit(0)
}

func printBranches() {
	branches := getBranchesWithDefaultIndication()

	for _, branch := range branches {
		fmt.Println(branch)
	}
}

func chooseBranch() string {
	branches := getBranchesWithDefaultIndication()
	cursor := 0

	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Press Enter to select a branch.")

		for i, branch := range branches {
			if i == cursor {
				fmt.Printf("> %s\n", branch)
			} else {
				fmt.Printf("  %s\n", branch)
			}
		}
	}
}

type branchChoice struct {
	branches        []string
	cursor          int
	selectedBranch  string
}

func initialModel() branchChoice {
	branches := getBranchesWithDefaultIndication()

	return branchChoice{
		branches:        branches,
		cursor:          0,
		selectedBranch:  "",
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
					menu.selectedBranch = menu.branches[menu.cursor]
					return menu, tea.Quit
			}
	}

	return menu, nil
}

func (m branchChoice) View() string {
	s := "\033[H\033[2J"
	s += "Choose a branch:\n\n"

	for i, branch := range m.branches {
		cursor := " "

		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, branch)
	}

	s += "\nPress q to quit.\n"

	return s
}