package libos

import (
	"fmt"
	"os"
)

func CreateFile(filename string, data []byte) error {
	// Membuka file dengan flag os.O_CREATE dan os.O_WRONLY
	// os.O_CREATE untuk membuat file jika belum ada
	// os.O_WRONLY untuk menulis ke dalam file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Menulis teks ke dalam file
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
