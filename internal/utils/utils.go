package utils

import (
	"bufio"
	"fmt"
	"os"
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