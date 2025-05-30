package builder

import (
	"fmt"
	"os"
)

func initializeDirProject(projectName string) error {
	_, err := os.ReadDir(projectName)
	if err == nil {
		return fmt.Errorf("project %s already exists", projectName)
	}

	if os.IsNotExist(err) {
		err := os.Mkdir(projectName, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating project %s: %w", projectName, err)
		}
	}

	return nil
}
