package os

import "os"

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) (*os.File, bool) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return &os.File{}, false
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return file, true
}
