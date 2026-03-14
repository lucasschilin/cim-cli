package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Confirm(question string) (bool, error) {
	tty, err := os.Open("/dev/tty")
	if err != nil {
		return false, nil
	}
	defer tty.Close()

	reader := bufio.NewReader(tty)

	fmt.Printf("%s (Y/n): ", question)

	answer, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	answer = strings.ToLower(strings.TrimSpace(answer))

	return (answer == "Y" || answer == "y" || answer == "yes"), nil
}
