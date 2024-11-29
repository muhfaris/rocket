package builder

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"os"
	"os/exec"

	"github.com/joncalhoun/pipe"
)

// CMDBuilder is wrap cmd data
type CMDBuilder struct {
	Project          string
	Pkg              string
	Name             string
	FormatType       string
	StorageType      string
	CacheType        string
	MessageQueueType string
}

// NewCMDBuilder is initialize cmd package
func NewCMDBuilder(project, storage, cache, queue string) *CMDBuilder {
	return &CMDBuilder{
		Project:          project,
		Pkg:              "cmd",
		Name:             "root",
		FormatType:       "toml",
		StorageType:      storage,
		CacheType:        cache,
		MessageQueueType: queue,
	}
}

// ChangeFormatType is change format type
func (cmd *CMDBuilder) ChangeFormatType(formatType string) {
	cmd.FormatType = formatType
}

func (cmd *CMDBuilder) IsRedis() bool {
	return cmd.CacheType == "redis"
}

func (cmd *CMDBuilder) IsPosgresSQL() bool {
	return cmd.StorageType == "postgresql"
}

func (cmd *CMDBuilder) generateData() map[string]interface{} {
	var isPostgresSQL, isRedis bool

	isPostgresSQL = cmd.IsPosgresSQL()
	isRedis = cmd.IsRedis()

	return map[string]interface{}{
		"pkg":           cmd.Pkg,
		"name":          cmd.Name,
		"isPostgresSQL": isPostgresSQL,
		"isRedis":       isRedis,
		"formatType":    cmd.FormatType,
		"chanSign":      template.HTML("<-"), // unescape string
	}
}

// Generate is generate cmd file
func (cmd *CMDBuilder) Generate() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// project name
	if _, err := os.Stat(cmd.Project); os.IsNotExist(err) {
		os.Mkdir(cmd.Project, 0o755)
	}

	pathCMD := fmt.Sprintf("%s/%s", cmd.Project, cmd.Pkg)
	if _, err := os.Stat(pathCMD); os.IsNotExist(err) {
		os.Mkdir(pathCMD, 0o755)
	}

	data := []generateTemplate{
		{
			project:      cmd.Project,
			name:         cmd.Name,
			pathMain:     pathCMD,
			pathTemplate: fmt.Sprintf("%s/builder/cmd/templates/root.tmpl", pwd),
			data:         cmd.generateData(),
		},
		{
			project:      cmd.Project,
			name:         "serve",
			pathMain:     pathCMD,
			pathTemplate: fmt.Sprintf("%s/builder/cmd/templates/serve.tmpl", pwd),
			data:         cmd.generateData(),
		},
	}

	return generateFiles(data)
}

type generateTemplate struct {
	project      string
	name         string
	pathMain     string
	pathTemplate string
	data         map[string]interface{}
}

func unescape(s string) template.HTML {
	return template.HTML(s)
}

func generateFiles(params []generateTemplate) error {
	for _, param := range params {
		t, err := template.ParseFiles(param.pathTemplate)
		if err != nil {
			return fmt.Errorf("error parse template cmd, %v", err)
		}

		rc, wc, _ := pipe.Commands(
			exec.Command("gofmt"),
			exec.Command("goimports"),
		)

		if err := t.Funcs(template.FuncMap{"unescape": unescape}).Execute(wc, param.data); err != nil {
			return fmt.Errorf("error funcmap unescape data when execute, %v", err)
		}

		wc.Close()
		var buf bytes.Buffer
		io.Copy(&buf, rc)

		p, err := format.Source(buf.Bytes())
		if err != nil {
			return fmt.Errorf("error format source data cmd, %v", err)
		}

		err = os.WriteFile(fmt.Sprintf("%s/%s.go", param.pathMain, param.name), p, 0o644)
		if err != nil {
			return fmt.Errorf("error re-write cmd file, %v", err)
		}
	}

	return nil
}
