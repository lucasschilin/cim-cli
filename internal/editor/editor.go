package editor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Open(path string) error {
	editor := detectEditor()
	if isInvalidEditor(editor) {
		return fmt.Errorf("Env dump: GIT_EDITOR=%q VISUAL=%q EDITOR=%q\n",
			os.Getenv("GIT_EDITOR"),
			os.Getenv("VISUAL"),
			os.Getenv("EDITOR"),
		)
	}

	parts := strings.Fields(editor)
	cmd := exec.Command(parts[0], append(parts[1:], path)...)

	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	defer tty.Close()

	cmd.Stdin = tty
	cmd.Stdout = tty
	cmd.Stderr = tty

	return cmd.Run()
}

func OpenTempFile() (string, error) {
	file, err := os.CreateTemp("", "cim-cli-commit-*.txt")
	if err != nil {
		return "", err
	}

	path := file.Name()
	file.Close()
	defer os.Remove(path)

	editor := detectEditor()
	if isInvalidEditor(editor) {
		return "", fmt.Errorf("no valid editor configured")
	}

	parts := strings.Fields(editor)
	editCmd := exec.Command(parts[0], append(parts[1:], path)...)

	editCmd.Stdin = os.Stdin
	editCmd.Stdout = os.Stdout
	editCmd.Stderr = os.Stderr

	if err := editCmd.Run(); err != nil {
		return "", err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func detectEditor() string {

	// variável do git
	if e := os.Getenv("GIT_EDITOR"); !isInvalidEditor(e) {
		return e
	}

	// configuração do git
	cmd := exec.Command("git", "config", "--get", "core.editor")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err == nil {
		editor := strings.TrimSpace(out.String())
		if editor != "" {
			return editor
		}
	}

	// VISUAL
	if e := os.Getenv("VISUAL"); !isInvalidEditor(e) {
		return e
	}

	// EDITOR
	if e := os.Getenv("EDITOR"); !isInvalidEditor(e) {
		return e
	}

	// fallback por sistema
	switch runtime.GOOS {
	case "windows":
		return "notepad"
	case "darwin":
		return "nano"
	default:
		return "nano"
	}
}

func isInvalidEditor(e string) bool {
	e = strings.TrimSpace(e)
	return e == "" || e == ":" || e == "true"
}
