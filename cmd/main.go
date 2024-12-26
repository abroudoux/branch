package main

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	root "github.com/abroudoux/branch"
	repository "github.com/abroudoux/branch/internal"
	tea "github.com/charmbracelet/bubbletea"
)

//go:embed ../config/config.json
var embeddedConfigFile []byte

var config Config

func main() {
	err := loadConfig()
	if err != nil {
		printErrorAndExit(err)
	}

	if len(os.Args) > 1 {
		err := flagMode()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	err = isGitInstalled()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = isInGitRepository()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	branch, err := chooseBranch()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	action, err := chooseAction(branch)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = doAction(branch, action)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Config struct {
	Ui struct {
		CursorColor string `json:"cursorColor"`
		BranchColor string `json:"branchColor"`
		BranchSelectedColor string `json:"branchSelectedColor"`
		ActionSelectedColor string `json:"actionSelectedColor"`
	} `json:"Ui"`
}

func loadConfig() error {
	err := json.Unmarshal(embeddedConfigFile, &config)
	if err != nil {
		return fmt.Errorf("error decoding embedded config: %v", err)
	}
	return nil
}

func isGitInstalled() error {
	cmd := exec.Command("git", "version")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("git is not installed: %v", err)
	}

	return nil
}

func isInGitRepository() error {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error checking if in git repository: %v", err)
	}

	return nil
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
	fmt.Printf("  %-20s %s\n", "branch [run | -r]", "Start the interactive branch selection")
	fmt.Printf("  %-20s %s\n", "branch [--list | -l]", "List all branches")
	fmt.Printf("  %-20s %s\n", "branch [--help | -h]", "Show this help message")
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

type BranchChoice struct {
	branches        []string
	cursor          int
	selectedBranch  string
}

func initialBranchModel() BranchChoice {
	branches := getBranchesWithDefaultIndication()

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

func chooseBranch() (string, error) {
	branchesMenu := tea.NewProgram(initialBranchModel())
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
		chooseBranch()
	case "-v", "--version":
		fmt.Println(root.AsciiArt)
		latestVersion, err := repository.GetLatestRelease()
		if err != nil {
			return fmt.Errorf("error getting latest version: %v", err)
		}

		fmt.Printf("Latest version: %s\n", latestVersion)
	case "-l", "--list":
		printBranches()
	case "-h", "--help":
		printHelpManual()
	default:
		printHelpManual()
	}

	return nil
}

func doAction(branch string, action string) error {
	switch action {
	case "Exit":
		fmt.Println("Exiting...")
		return nil
	case "Delete":
		return deleteBranch(branch)
	case "Merge":
		return mergeBranch(branch)
	case "Branch":
		return createBranch(branch)
	case "Rename":
		return renameBranch(branch)
	case "Checkout":
		return checkoutBranch(branch)
	case "Name":
		return copyName(branch)
	default:
		return fmt.Errorf("invalid action: %s", action)
	}
}

func deleteBranch(branch string) error {
	if !askConfirmation(fmt.Sprintf("Are you sure you want to delete '%s'?", renderBranch(branch))) {
		return fmt.Errorf("branch deletion cancelled")
	}

	cmd := exec.Command("git", "branch", "-D", branch)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error deleting branch: %v", err)
	}

	println("Branch '%s' deleted", renderBranch(branch))

	if hasRemoteBranch(branch) && askConfirmation(fmt.Sprintf("Do you want to delete '%s' remotly?", renderBranch(branch))) {
		err := deleteRemoteBranch(branch)
		if err != nil {
			return fmt.Errorf("error deleting remote branch: %v", err)
		}
	}
	return nil
}

func deleteRemoteBranch(branch string) error {
	if !askConfirmation(fmt.Sprintf("Are you sure you want to delete '%s' remotly?", renderBranch(branch))) {
		return fmt.Errorf("branch deletion cancelled")
	}

	cmd := exec.Command("git", "push", "origin", "--delete", branch)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error deleting remote branch: %v", err)
	}

	return nil
}

func hasRemoteBranch(branchName string) bool {
	cmd := exec.Command("git", "ls-remote", "--heads", "origin", branchName)
	err := cmd.Run()
	return err == nil
}

func mergeBranch(branch string) error {
	cmd := exec.Command("git", "merge", branch)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error merging branch", err)
		return fmt.Errorf("error merging branch: %v", err)
	}

	println("Branch '%s' merged", branch)

	shouldDeleteBranch := askConfirmation("Do you want to delete the merged branch?")
	if shouldDeleteBranch {
		err := deleteBranch(branch)
		if err != nil {
			return fmt.Errorf("error deleting branch: %v", err)
		}
	}

	println("Branch '%s' deleted", renderBranch(branch))

	return nil
}

func createBranch(branch string) error {
	newBranchName, err := askInput("Enter the name of the new branch: ")
	if err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}

	branches := getBranches()
	for _, branch := range branches {
		if branch == newBranchName {
			return fmt.Errorf("branch '%s' already exists", renderBranch(newBranchName))
		}
	}

	defaultBranch := getDefaultBranch()
	if branch != defaultBranch {
		cmd := exec.Command("git", "checkout", branch)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("error checking out default branch: %v", err)
		}
	}

	if (askConfirmation("Do you want to checkout on the new branch?")) {
		cmd := exec.Command("git", "checkout", "-b", newBranchName)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("error creating branch: %v", err)
		}
	} else {
		cmd := exec.Command("git", "branch", newBranchName)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("error creating branch: %v", err)
		}
	}

	fmt.Printf("Branch '%s' based on '%s' created\n", renderBranch(newBranchName), renderBranch(branch))
	return nil
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

func askInput(message string) (string, error) {
	var input string
	fmt.Print(message)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", fmt.Errorf("error reading input: %v", err)
	}

	return input, nil
}

func renameBranch(branch string) error {
	newBranchName, err := askInput("Enter a name for the new branch: ")
	if err != nil {
		return fmt.Errorf("error reading input: %v", err)
	}

	if strings.Contains(newBranchName, " ") {
		return fmt.Errorf("error: the branch name must not contain spaces")
	}

	cmd := exec.Command("git", "branch", "-m", branch, newBranchName)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error renaming branch: %v", err)
	}

	fmt.Printf("Branch %s renamed to %s\n", renderBranch(branch), renderBranch(newBranchName))
	return nil
}

func checkoutBranch(branch string) error {
	cmd := exec.Command("git", "checkout", branch)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error checking out branch: %v", err)
	}

	fmt.Printf("Branch %s checked out\n", renderBranch(branch))
	return nil
}

func copyName(branch string) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(branch)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error copying branch name: %v", err)
	}

	fmt.Printf("Branch name '%s' copied to clipboard\n", renderBranch(branch))
	return nil
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

func printErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}