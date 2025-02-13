package logs

import "github.com/charmbracelet/log"

func Info(msg string) {
	log.Info(msg)
}

func Warn(msg string, err error) {
	log.Warn(msg, err)
}

func WarnMsg(msg string) {
	log.Warn(msg, nil)
}

func Error(msg string, err error) {
	log.Error(msg, err)
}
