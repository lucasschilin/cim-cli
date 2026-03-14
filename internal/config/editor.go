package config

import (
	"os"
	"os/exec"
	"path/filepath"
)

func EnsureConfigFile(path string) error {

	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {

		file, err := os.Create(path)
		if err != nil {
			return err
		}

		defer file.Close()
	}

	return nil
}

// OpenEditor abre o arquivo de configuração usando o editor do sistema.
// Ele segue a ordem padrão de ferramentas Unix:
//
// 1. variável de ambiente EDITOR
// 2. variável VISUAL
// 3. fallback para "vi"
func OpenEditor(path string) error {

	editor := os.Getenv("EDITOR")

	if editor == "" {
		editor = os.Getenv("VISUAL")
	}

	if editor == "" {
		editor = "vi"
	}

	cmd := exec.Command(editor, path)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
