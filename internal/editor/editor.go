package editor

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func Open(path string) error {
	editor := detectEditor()

	cmd := exec.Command(editor, path)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func detectEditor() string {

	// variável do git
	if e := os.Getenv("GIT_EDITOR"); e != "" {
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
	if e := os.Getenv("VISUAL"); e != "" {
		return e
	}

	// EDITOR
	if e := os.Getenv("EDITOR"); e != "" {
		return e
	}

	// fallback
	return "vi"
}
