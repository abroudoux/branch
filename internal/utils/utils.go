package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func AskConfirmation(message string) bool {
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

func AskInput(message string) (string, error) {
	var input string
	fmt.Print(message)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", fmt.Errorf("error reading input: %v", err)
	}

	return input, nil
}

func PrintErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func CleanString(s string) string {
	return strings.TrimSpace(strings.TrimPrefix(s, "*"))
}

func isGitInstalled() error {
	cmd := exec.Command("git", "version")
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func IsInGitRepository() error {
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

func PrintBranches(branches []string) {
	for _, branch := range branches {
		fmt.Println(branch)
	}
}