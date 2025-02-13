package forms

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskConfirmation(message string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (y/n) [yes]: ", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	confirmation := strings.TrimSpace(input)
	if confirmation == "" || strings.EqualFold(confirmation, "y") || strings.EqualFold(confirmation, "yes") {
		return true, nil
	}

	return false, nil

	return true, nil
}

func AskInput(message string) (string, error) {
	var input string

	fmt.Print(message)

	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}

	return input, nil
}
