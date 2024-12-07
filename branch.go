package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]

		if arg == "-v" || arg == "--verbose" {
			fmt.Println("2.0.0")
		} else if arg == "-l" || arg == "--list" {
			printBranches()
		} else if arg == "-h" || arg == "--help" {
			printHelpManual()
			os.Exit(0)
		}
	}

	isGitInstalled()
	isInGitRepository()

	//branches := getBranchesWithDefaultIndication()
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
}

func printBranches() {
	branches := getBranchesWithDefaultIndication()

	for _, branch := range branches {
		fmt.Println(branch)
	}
}