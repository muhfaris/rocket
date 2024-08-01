package builder

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/getkin/kin-openapi/openapi3"
	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
)

var _baseproject BaseProject

type Main struct {
	doc               *openapi3.T
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

func New(doc *openapi3.T, packagePath, projectName string) *Main {
	_baseproject = BaseProject{
		AppName:     projectName,
		ProjectName: projectName,
		PackagePath: packagePath,
	}

	return &Main{
		doc:               doc,
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

	// slog.Info("Creating new project", "project", _baseproject.ProjectName)
	fmt.Println("Creating new project", _baseproject.ProjectName)

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
	// slog.Info("└── Creating .gitignore", "project", _baseproject.ProjectName)
	fmt.Println("└── Creating .gitignore", _baseproject.ProjectName)
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
	// slog.Info("└── Formatting directory", "project", _baseproject.ProjectName)
	fmt.Println("└── Formatting directory", _baseproject.ProjectName)
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

	project := NewProject(m.doc, _baseproject.ProjectName)
	err = project.GenerateDirectories()
	if err != nil {
		return err
	}

	err = m.initializeModule()
	if err != nil {
		return err
	}

	return nil
}

func (m *Main) initializeModule() error {
	// slog.Info("└── Initializing go module", "project", _baseproject.ProjectName)
	fmt.Println("└── Initializing go module", _baseproject.ProjectName)
	// Change the current working directory to the specific directory
	err := os.Chdir(_baseproject.ProjectName)
	if err != nil {
		return fmt.Errorf("failed to change directory to %s: %v", _baseproject.ProjectName, err)
	}

	// Initialize the Go module
	err = exec.Command("go", "mod", "init", _baseproject.PackagePath).Run()
	if err != nil {
		return fmt.Errorf("failed to initialize go module: %v", err)
	}

	err = exec.Command("go", "mod", "tidy").Run()
	if err != nil {
		return fmt.Errorf("failed to tidy go module: %v", err)
	}

	err = exec.Command("go", "mod", "vendor").Run()
	if err != nil {
		return fmt.Errorf("failed to vendor go module: %v", err)
	}

	return nil
}
