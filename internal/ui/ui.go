package ui

import "fmt"

func RenderCursor() string {
	render := fmt.Sprintf("\033[%sm>\033[0m", "32")
	return render
}

func RenderBranch(branchName string) string {
    return fmt.Sprintf("\033[%sm%s\033[0m", "38;2;214;112;214", branchName)
}

func RenderBranchSelected(branchName string, isSelected bool) string {
    if isSelected {
        return fmt.Sprintf("\033[%sm%s\033[0m", "32", branchName)
    }
    return branchName
}

func RenderActionSelected(action string, isSelected bool) string {
    if isSelected {
        return fmt.Sprintf("\033[%sm%s\033[0m", "32", action)
    }
    return action
}