package repository

import (
	"fmt"
	"net/http"
	"os"
)

func PrintHelpManual() {
	fmt.Println("Usage: branch [options]")
	fmt.Printf("  %-20s %s\n", "branch [--help | -h]", "Show this help message")
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
	flag := os.Args[1]

	switch flag {
	case "-v", "--version":
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