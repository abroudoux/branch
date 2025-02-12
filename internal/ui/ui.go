package ui

import "fmt"

func RenderElementSelected(el string) string {
	return fmt.Sprintf("\033[%sm%s\033[0m", "38;2;214;112;214", el)
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

func RenderCursor(isCurrentLine bool) string {
	if isCurrentLine {
		return fmt.Sprintf("\033[%sm>\033[0m", "32")
	}

	return " "
}

func RenderCurrentLine(s string, isCurrentLine bool) string {
	if isCurrentLine {
		return fmt.Sprintf("\033[%sm%s\033[0m", "32", s)
	}
	return s
}
