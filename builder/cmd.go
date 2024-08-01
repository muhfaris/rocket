package builder

import (
	"fmt"
	"os"

	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
)

type CMD struct {
	template []byte
	dirpath  string
	filepath string
	AppName  string
}

type CMDData struct {
	AppName     string
	PackagePath string
}

func newCMD(appName, projectName string) *CMD {
	return &CMD{
		template: templates.GetCMDTemplate(),
		dirpath:  fmt.Sprintf("%s/cmd", projectName),
		filepath: fmt.Sprintf("%s/cmd/root.go", projectName),
		AppName:  appName,
	}
}

func (c *CMD) Generate() error {
	// slog.Info("└── Creating cmd")
	fmt.Println("└── Creating cmd")

	data := CMDData{
		AppName:     c.AppName,
		PackagePath: _baseproject.PackagePath,
	}

	_, err := os.Stat(c.dirpath)
	if os.IsExist(err) {
		return fmt.Errorf("directory project %s already exists", c.dirpath)
	}

	if os.IsNotExist(err) {
		err := os.Mkdir(c.dirpath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory cmd %s: %w", c.dirpath, err)
		}
	}

	raw, err := libos.ExecuteTemplate(c.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(c.filepath, raw)
	if err != nil {
		return err
	}

	return nil
}
