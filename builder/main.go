package builder

import (
	"fmt"

	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
)

var _baseproject BaseProject

type Main struct {
	template          []byte
	tempalteGitignore []byte
	filename          string
	filepath          string
	BasePackage
}

type MainData struct {
	PackagePath string
	ProjectName string
	Path        string
}

func New(packagePath, projectName string) *Main {
	_baseproject = BaseProject{
		AppName:     projectName,
		ProjectName: projectName,
		PackagePath: packagePath,
	}
	return &Main{
		template:          templates.GetMainTemplate(),
		tempalteGitignore: templates.GetGitIgnore(),
		filename:          "main.go",
		filepath:          fmt.Sprintf("%s/main.go", projectName),
		BasePackage: BasePackage{
			PackageName: "main",
			PackagePath: packagePath,
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
		err = libos.DeleteDir(_baseproject.ProjectName)
		if err != nil {
			return
		}
	}()

	// create project
	err = initializeDirProject(_baseproject.ProjectName)
	if err != nil {
		return err
	}

	// create main.go
	data := MainData{
		PackagePath: m.PackagePath,
		Path:        "cmd",
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
	filepathGitignore := fmt.Sprintf("%s/.gitignore", _baseproject.ProjectName)
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
	err = libos.FormatDirPath(_baseproject.ProjectName)
	if err == nil {
		return nil
	}

	return nil
}

func (m *Main) generate() error {
	cmd := newCMD(_baseproject.AppName, _baseproject.ProjectName)
	err := cmd.Generate()
	if err != nil {
		return err
	}

	cfg := NewConfig("config", "yaml", _baseproject.ProjectName)
	err = cfg.Generate()
	if err != nil {
		return err
	}

	return nil
}
