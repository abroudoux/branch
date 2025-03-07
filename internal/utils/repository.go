package utils

import (
	"fmt"
)

func getLatestVersion() string {
	return "v0.1.2"
}

func Version() {
	printAscii()
	fmt.Printf("Latest version: %s\n", getLatestVersion())
}

func Help() {
	commands := []string{
		"branch",
		"branch [--help, -h]",
	}
	descriptions := []string{
		"Run the program",
		"Show this help message",
	}

	fmt.Println("\nUsage: branch [options]")
	for i, cmd := range commands {
		fmt.Printf("  %-20s %s\n", cmd, descriptions[i])
	}
}
