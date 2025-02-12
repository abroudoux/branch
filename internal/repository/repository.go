package repository

import (
	"fmt"
	"net/http"
	"os"

	"github.com/abroudoux/branch/internal/ascii"
)

func PrintHelpManual() {
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

func GetLatestRelease() (string, error) {
	url := "https://api.github.com/repos/abroudoux/branch/releases/latest"
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	latestVersion := res.Header.Get("tag_name")
	return latestVersion, nil
}

func FlagMode() error {
	otpion := os.Args[1]

	switch otpion {
	case "-v", "--version":
		err := ascii.PrintAscii()
		if err != nil {
			return err
		}

		latestVersion, err := GetLatestRelease()
		if err != nil {
			return err
		}

		fmt.Printf("Latest version: %s\n", latestVersion)
	case "-h", "--help":
		PrintHelpManual()
	default:
		PrintHelpManual()
	}

	return nil
}
