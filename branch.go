package main

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

//go:embed assets/ascii.txt
var asciiArt string

type Config struct {
	Ui struct {
		CursorColor string `json:"cursorColor"`
		BranchColor string `json:"branchColor"`
		BranchSelectedColor string `json:"branchSelectedColor"`
		ActionSelectedColor string `json:"actionSelectedColor"`
	} `json:"Ui"`
}

var config Config

func loadConfig() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error loading config file:", err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		os.Exit(1)
	}
}

func main() {
	loadConfig()

	if len(os.Args) > 1 {
		flagMode()
	}

	isGitInstalled()
	isInGitRepository()

	branch := chooseBranch()

	if branch == "" {
		fmt.Println("No branch selected. Exiting...")
		return
	}

	action := chooseAction(branch)

	if action == "" {
		fmt.Println("No action selected. Exiting...")
		return
	}

	doAction(branch, action)
}

func isGitInstalled() {
	cmd := exec.Command("git", "version")
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error checking if git is installed", err)
		os.Exit(1)
	}
}

func isInGitRepository() {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error checking if in git repository", err)
		os.Exit(1)
	}
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
			branchesWithDefaultIndication = append(branchesWithDefaultIndication, "* "+branch)
		} else {
			branchesWithDefaultIndication = append(branchesWithDefaultIndication, "  "+branch)
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

func cleanString(s string) string {
	return strings.TrimSpace(strings.TrimPrefix(s, "*"))
}

type branchChoice struct {
	branches        []string
	cursor          int
	selectedBranch  string
}

func initialBranchModel() branchChoice {
	branches := getBranchesWithDefaultIndication()

	return branchChoice{
		branches:        branches,
		cursor:          len(branches) - 1,
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

func (menu branchChoice) View() string {
    s := "\033[H\033[2J"
    s += "Choose a branch:\n\n"

    for i, branch := range menu.branches {
        cursor := " "

        if menu.cursor == i {
            cursor = renderCursor()
            s += fmt.Sprintf("%s %s\n", cursor, renderBranchSelected(branch, true))
        } else {
            s += fmt.Sprintf("%s %s\n", cursor, renderBranchSelected(branch, false))
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
    s += fmt.Sprintf("Branch: %s\n\n", renderBranch(menu.selectedBranch))
	s += "Choose an action:\n\n"

	for i, action := range menu.actions {
		cursor := " "

		if menu.cursor == i {
            cursor = renderCursor()
            s += fmt.Sprintf("%s %s\n", cursor, renderActionSelected(action, true))
        } else {
            s += fmt.Sprintf("%s %s\n", cursor, renderActionSelected(action, false))
        }
	}

	return s
}

func chooseBranch() string {
	branchesMenu := tea.NewProgram(initialBranchModel())
	finalModel, err := branchesMenu.Run()

	if err != nil {
		fmt.Printf("Error running the interactive menu: %v\n", err)
		os.Exit(1)
	}

	branchMenu := finalModel.(branchChoice)
	return cleanString(branchMenu.selectedBranch)
}

func chooseAction(selectedBranch string) string {
	actionsMenu := tea.NewProgram(initialActionModel(selectedBranch))
	finalActionModel, err := actionsMenu.Run()

	if err != nil {
		fmt.Printf("Error running the actions menu: %v\n", err)
		os.Exit(1)
	}

	actionMenu := finalActionModel.(actionChoice)
	return cleanString(actionMenu.selectedAction)
}

func flagMode() {
	arg := os.Args[1]

	if arg == "run" || arg == "-r" {
		chooseBranch()
	} else if arg == "-v" || arg == "--verbose" {
		fmt.Println(asciiArt)
		fmt.Println("2.0.1")
	} else if arg == "-l" || arg == "--list" {
		printBranches()
	} else if arg == "-h" || arg == "--help" {
		printHelpManual()
	}

	os.Exit(0)
}

func doAction(branch string, action string) {
	switch action {
		case "Exit":
			fmt.Println("Exiting...")
			return
		case "Delete":
			deleteBranch(branch)
			return
		case "Merge":
			mergeBranch(branch)
			return
		case "Branch":
			createBranch(branch)
			return
		case "Rename":
			renameBranch(branch)
			return
		case "Checkout":
			checkoutBranch(branch)
			return
		case "Name":
			copyName(branch)
			return
	}
}

func deleteBranch(branch string) {
	if !askConfirmation(fmt.Sprintf("Are you sure you want to delete '%s'?", renderBranch(branch))) {
		fmt.Println("Branch deletion cancelled")
		return
	}

	cmd := exec.Command("git", "branch", "-D", branch)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error deleting branch", err)
		os.Exit(1)
	}

	fmt.Printf("Branch %s deleted\n", renderBranch(branch))
}

func mergeBranch(branch string) {
	cmd := exec.Command("git", "merge", branch)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error merging branch", err)
		os.Exit(1)
	}

	fmt.Printf("Branch %s merged\n", branch)
}

func createBranch(branch string) {
	newBranchName := askInput("Enter the name of the new branch: ")

	branches := getBranches()
	for _, branch := range branches {
		if branch == newBranchName {
			fmt.Printf("Branch '%s' already exists.\n", renderBranch(newBranchName))
			return
		}
	}

	defaultBranch := getDefaultBranch()
	if branch != defaultBranch {
		cmd := exec.Command("git", "checkout", branch)
		err := cmd.Run()

		if err != nil {
			fmt.Println("Error checking out default branch:", err)
			os.Exit(1)
		}
	}

	if (askConfirmation("Do you want to checkout on the new branch?")) {
		cmd := exec.Command("git", "checkout", "-b", newBranchName)
		err := cmd.Run()

		if err != nil {
			fmt.Println("Error creating branch:", err)
			os.Exit(1)
		}
	} else {
		cmd := exec.Command("git", "branch", newBranchName)
		err := cmd.Run()

		if err != nil {
			fmt.Println("Error creating branch:", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Branch '%s' based on '%s' created\n", renderBranch(newBranchName), renderBranch(branch))
}

func askConfirmation(message string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (y/n) [yes]: ", message)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	confirmation := strings.TrimSpace(input)

	if confirmation == "" || strings.EqualFold(confirmation, "y") || strings.EqualFold(confirmation, "yes") {
		return true
	}

	return false
}

func askInput(message string) string {
	var input string
	fmt.Print(message)
	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	return input
}

func renameBranch(branch string) {
	newBranchName := askInput("Enter the new name for the branch: ")

	cmd := exec.Command("git", "branch", "-m", branch, newBranchName)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error renaming branch", err)
		os.Exit(1)
	}

	fmt.Printf("Branch %s renamed to %s\n", renderBranch(branch), renderBranch(newBranchName))
}

func checkoutBranch(branch string) {
	cmd := exec.Command("git", "checkout", branch)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error checking out branch", err)
		os.Exit(1)
	}

	fmt.Printf("Branch %s checked out\n", renderBranch(branch))
}

func copyName(branch string) {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(branch)
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error copying branch name:", err)
		os.Exit(1)
	}

	fmt.Printf("Branch name '%s' copied to clipboard\n", renderBranch(branch))
}

func renderCursor() string {
	render := fmt.Sprintf("\033[%sm>\033[0m", config.Ui.CursorColor)
	return render
}

func renderBranch(branchName string) string {
    return fmt.Sprintf("\033[%sm%s\033[0m", config.Ui.BranchColor, branchName)
}

func renderBranchSelected(branchName string, isSelected bool) string {
    if isSelected {
        return fmt.Sprintf("\033[%sm%s\033[0m", config.Ui.BranchSelectedColor, branchName)
    }
    return branchName
}

func renderActionSelected(action string, isSelected bool) string {
    if isSelected {
        return fmt.Sprintf("\033[%sm%s\033[0m", config.Ui.ActionSelectedColor, action)
    }
    return action
}