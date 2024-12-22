package builder

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/muhfaris/rocket/builder/hexagonal"
	libos "github.com/muhfaris/rocket/shared/os"
	"github.com/muhfaris/rocket/shared/templates"
)

var _baseproject BaseProject

type Main struct {
	arch              string
	content           []byte
	doc               *openapi3.T
	template          []byte
	tempalteGitignore []byte
	filename          string
	filepath          string
	cacheType         string
	dbType            string
	BasePackage
}

type MainData struct {
	PackagePath string
	ProjectName string
	Path        string
}

func New(content []byte, doc *openapi3.T, packagePath, projectName, arch, cacheParam, dbParam string) *Main {
	_baseproject = BaseProject{
		AppName:     projectName,
		ProjectName: projectName,
		PackagePath: packagePath,
	}

	return &Main{
		arch:              arch,
		content:           content,
		doc:               doc,
		template:          templates.GetMainTemplate(),
		tempalteGitignore: templates.GetGitIgnore(),
		filename:          "main.go",
		filepath:          fmt.Sprintf("%s/main.go", projectName),
		cacheType:         cacheParam,
		dbType:            dbParam,
		BasePackage: BasePackage{
			PackageName: "main",
			PackagePath: packagePath,
		},
	}
}

func (m *Main) Generate() error {
	var err error
	// slog.Info("Creating new project", "project", _baseproject.ProjectName)
	fmt.Println("Creating new project", _baseproject.ProjectName)

	// create project
	err = initializeDirProject(_baseproject.ProjectName)
	if err != nil {
		return err
	}

	err = m.CreateOASSpec()
	if err != nil {
		return err
	}

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

	err = m.initMain()
	if err != nil {
		return err
	}

	// create .gitignore
	err = m.initGitignore()
	if err != nil {
		return err
	}

	// all code files will Generate
	err = m.generate()
	if err != nil {
		return err
	}

	// format file go
	fmt.Println("└── Formatting directory")
	time.Sleep(10 * time.Millisecond)
	fmt.Printf(" %s goimports\n", LineOnProgress)
	err = m.GoImports(_baseproject.ProjectName)
	if err != nil {
		return err
	}

	fmt.Printf(" %s gofmt\n", LineLast)
	err = m.GoFmt(_baseproject.ProjectName)
	if err != nil {
		return err
	}

	return nil
}

func (m *Main) generate() error {
	cmd := newCMD(_baseproject.AppName, _baseproject.ProjectName)
	err := cmd.Generate()
	if err != nil {
		return err
	}

	cfg := NewConfig("config", "yaml", _baseproject.ProjectName, m.cacheType, m.dbType)
	err = cfg.Generate()
	if err != nil {
		return err
	}

	switch m.arch {
	case "hexagonal":
		based := hexagonal.Based{
			Project: hexagonal.BaseProject{
				AppName:     _baseproject.AppName,
				ProjectName: _baseproject.ProjectName,
				PackagePath: _baseproject.PackagePath,
			},
			Package: hexagonal.BasePackage{},
		}

		project := hexagonal.NewProject(m.doc, based, _baseproject.ProjectName, m.cacheType)
		project.GenerateDirectories()
	}

	// project := NewProject(m.doc, _baseproject.ProjectName, m.cacheType)
	// err = project.GenerateDirectories()
	// if err != nil {
	// 	return err
	// }

	err = m.initializeModule()
	if err != nil {
		return err
	}

	return nil
}

func (m *Main) initializeModule() error {
	fmt.Printf("%s%s\n", LineLast, "Go module")
	// Initialize the Go module
	cmd := exec.Command("go", "mod", "init", _baseproject.PackagePath)
	cmd.Dir = _baseproject.ProjectName
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to initialize go module: %v", err)
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = _baseproject.ProjectName
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to tidy go module: %v", err)
	}

	cmd = exec.Command("go", "mod", "vendor")
	cmd.Dir = _baseproject.ProjectName
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to vendor go module: %v", err)
	}

	return nil
}

func (m *Main) initMain() error {
	fmt.Printf("%s%s\n", LineOnProgress, m.filename)
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

	return nil
}

func (m *Main) initGitignore() error {
	fmt.Printf("%s%s\n", LineOnProgress, ".gitignore")
	filepathGitignore := fmt.Sprintf("%s/.gitignore", _baseproject.ProjectName)
	err := libos.CreateFile(filepathGitignore, m.tempalteGitignore)
	if err != nil {
		return err
	}

	return nil
}

func (m *Main) GoImports(directory string) error {
	cmd := exec.Command("goimports", "-w", directory)
	return cmd.Run()
}

func (m *Main) GoFmt(directory string) error {
	cmd := exec.Command("gofmt", "-w", directory)
	return cmd.Run()
}

func (m *Main) CreateOASSpec() error {
	dirpath := fmt.Sprintf("%s/spec", _baseproject.ProjectName)
	err := os.MkdirAll(dirpath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating dir %s: %w", dirpath, err)
	}

	filepath := fmt.Sprintf("%s/openapi.yaml", dirpath)
	err = libos.CreateFile(filepath, m.content)
	if err != nil {
		return fmt.Errorf("error creating file %s: %w", m.filename, err)
	}
	return nil
}
