package repository

import "fmt"

func PrintHelpManual() {
	fmt.Println("Usage: branch [options]")
	fmt.Printf("  %-20s %s\n", "branch [run | -r]", "Start the interactive branch selection")
	fmt.Printf("  %-20s %s\n", "branch [--list | -l]", "List all branches")
	fmt.Printf("  %-20s %s\n", "branch [--help | -h]", "Show this help message")
}