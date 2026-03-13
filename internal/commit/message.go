package commit

import (
	"os"
	"strings"
)

func ReadCommitMessage(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	message := string(data)

	message = strings.TrimSpace(message)

	return message, nil
}

func WriteCommitMessage(path string, message string) error {

	data := []byte(message)

	return os.WriteFile(path, data, 0644)
}
