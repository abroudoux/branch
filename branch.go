package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	isGitInstalled()
	isInGitRepository()

	branches := getBranches()

	if len(branches) == 0 {
		fmt.Println("No branches found")
		os.Exit(1)
	}

	fmt.Println("Branches:")
	for _, branch := range branches {
		fmt.Println(branch)
	}
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
	cmd := exec.Command("git", "branch", "--format='%(refname:short)'")
	branches, err := cmd.Output()

	if err != nil {
		fmt.Println("Error getting branches", err)
		os.Exit(1)
	}

	return strings.Fields(string(branches))
}