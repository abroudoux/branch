package ui

import (
	"fmt"
)

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

func RenderCurrentBranch(s string, isCurrentLine bool, isHead bool) string {
	if isCurrentLine {
		if isHead {
			return fmt.Sprintf("\033[%sm%s\033[0m", "32", s)
		} else {
			return fmt.Sprintf("\033[%sm%s\033[0m", "32", s)
		}
	}
	return s
}

func RenderElementSelected(el string) string {
	return fmt.Sprintf("\033[%sm%s\033[0m", "38;2;214;112;214", el)
}
