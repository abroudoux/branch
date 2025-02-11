package ascii

import (
	"fmt"
	"os"
)

func PrintAscii() error {
	ascii, err := os.ReadFile("./ressources/ascii.txt")
    if err != nil {
		return err
    }

    fmt.Println(string(ascii))
	return nil
}