package internal

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

func HandleFlags() {
	if len(os.Args) > 1 {
		option := os.Args[1]

		switch option {
		case "--help", "-h":
			printHelpManual()
		case "--version", "-v":
			printLastVersion()
		default:
			log.Warn("Unknown option.")
			printHelpManual()
		}

		os.Exit(0)
	}
}

func getLatestVersion() string {
	return "v0.1.2"
}

func printLastVersion() {
	printAscii()
	fmt.Printf("Latest version: %s\n", getLatestVersion())
}

func printHelpManual() {
	cmds := []string{
		"branch",
		"branch [--help, -h]",
	}
	descs := []string{
		"Run the program",
		"Show this help message",
	}

	fmt.Println("\nUsage: branch [options]")
	for i, cmd := range cmds {
		fmt.Printf("  %-20s %s\n", cmd, descs[i])
	}
}

func printAscii() {
	ascii, _ := os.ReadFile("./ressources/ascii.txt")
	fmt.Println(string(ascii))
}
