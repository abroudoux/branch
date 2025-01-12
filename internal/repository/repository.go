package repository

import (
	"fmt"
	"net/http"
)

func PrintHelpManual() {
	fmt.Println("Usage: branch [options]")
	fmt.Printf("  %-20s %s\n", "branch [run | -r]", "Start the interactive branch selection")
	fmt.Printf("  %-20s %s\n", "branch [--list | -l]", "List all branches")
	fmt.Printf("  %-20s %s\n", "branch [--help | -h]", "Show this help message")
}

func GetLatestRelease() (string, error) {
	url := "https://api.github.com/repos/abroudoux/branch/releases/latest"
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error while fetching latest release: %v", err)
	}

	latestVersion := res.Header.Get("tag_name")
	return latestVersion, nil
}