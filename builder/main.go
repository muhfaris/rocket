package builder

import (
	"fmt"

	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
)

type Main struct {
	template          []byte
	tempalteGitignore []byte
	filename          string
	filepath          string
	ProjectName       string
	Base
}

type MainData struct {
	PackagePath string
	Path        string
}

func New(packagePath, projectName string) *Main {
	return &Main{
		template:          templates.GetMainTemplate(),
		tempalteGitignore: templates.GetGitIgnore(),
		filename:          "main.go",
		filepath:          fmt.Sprintf("%s/main.go", projectName),
		ProjectName:       projectName,
		Base: Base{
			PackagePath: packagePath,
			Path:        "cmd",
		},
	}
}

func (m *Main) Generate() error {
	var err error
	defer func() {
		if err == nil {
			return
		}
		// failed to create new project
		// delete created project
		err = libos.DeleteDir(m.ProjectName)
		if err != nil {
			return
		}
	}()

	// create project
	err = initializeDirProject(m.ProjectName)
	if err != nil {
		return err
	}

	// create main.go
	data := MainData{
		PackagePath: m.PackagePath,
		Path:        m.Path,
	}

	raw, err := libos.ExecuteTemplate(m.template, data)
	if err != nil {
		return err
	}

	err = libos.CreateFile(m.filepath, raw)
	if err != nil {
		return err
	}

	// create .gitignore
	filepathGitignore := fmt.Sprintf("%s/.gitignore", m.ProjectName)
	err = libos.CreateFile(filepathGitignore, m.tempalteGitignore)
	if err != nil {
		return err
	}

	// all code files will Generate
	err = m.generate()
	if err != nil {
		return err
	}

	// format file go
	err = libos.FormatDirPath(m.ProjectName)
	if err == nil {
		return nil
	}

	return nil
}

func (m *Main) generate() error {
	cmd := newCMD(m.ProjectName, m.ProjectName)
	err := cmd.Generate()
	if err != nil {
		return err
	}

	return nil
}
