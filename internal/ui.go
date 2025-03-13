package internal

import "fmt"

func renderCursor(isCurrentLine bool) string {
	if isCurrentLine {
		return fmt.Sprintf("\033[%sm>\033[0m", "32")
	}

	return " "
}
func renderCurrentLine(s string, isCurrentLine bool) string {
	if isCurrentLine {
		return fmt.Sprintf("\033[%sm%s\033[0m", "32", s)
	}
	return s
}

func renderElSelected(el string) string {
	return fmt.Sprintf("\033[%sm%s\033[0m", "38;2;214;112;214", el)
}
