package utils

import (
	"fmt"
	"os"
)

func printAscii() {
	ascii, _ := os.ReadFile("./ressources/ascii.txt")
	fmt.Println(string(ascii))
}
