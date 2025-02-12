package utils

import (
	"fmt"
	"os"
)

func getLatestVersion() string {
	return "v0.1.2"
}

func printHelpManual() {
	commands := []string{
		"branch",
		"branch [--help, -h]",
	}
	descriptions := []string{
		"Run the program",
		"Show this help message",
	}

	fmt.Println("Usage: branch [options]")
	for i, cmd := range commands {
		fmt.Printf("  %-20s %s\n", cmd, descriptions[i])
	}
}

func FlagMode() error {
	option := os.Args[1]

	switch option {
	case "-v", "--version":
		printAscii()
		fmt.Printf("Latest version: %s\n", getLatestVersion())
	case "-h", "--help":
		printHelpManual()
	default:
		printHelpManual()
	}

	return nil
}
