package libos

import (
	"fmt"
	"os"
	"os/exec"
)

func FormatDirPath(dirPath string) error {
	// Menyiapkan perintah gofmt -w .
	cmd := exec.Command("gofmt", "-w", dirPath)

	// Menjalankan perintah
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error formatting directory %s: %w", dirPath, err)
	}

	return nil
}

func DeleteDir(dirPath string) error {
	err := os.RemoveAll(dirPath)
	if err != nil {
		return fmt.Errorf("error deleting directory %s: %w", dirPath, err)
	}

	return nil
}
