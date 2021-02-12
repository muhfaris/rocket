package builder

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type cmdBuilder struct {
	project          string
	pkg              string
	name             string
	formatType       string
	storageType      string
	cacheType        string
	messageQueueType string
}

func NewCMDBuilder() *cmdBuilder {
	return &cmdBuilder{
		project:          "rocket-sample",
		pkg:              "cmd",
		name:             "root",
		formatType:       "toml",
		storageType:      "postgresql",
		cacheType:        "none",
		messageQueueType: "none",
	}
}

// ChangeFormatType is change format type
func (cmd *cmdBuilder) ChangeFormatType(formatType string) {
	cmd.formatType = formatType
}

func (cmd *cmdBuilder) IsRedis() bool {
	return cmd.cacheType == "redis"
}

func (cmd *cmdBuilder) IsPosgresSQL() bool {
	return cmd.storageType == "postgresql"
}

func (cmd *cmdBuilder) generateData() map[string]interface{} {
	var isPostgresSQL, isRedis bool

	isPostgresSQL = cmd.IsPosgresSQL()
	isRedis = cmd.IsRedis()

	return map[string]interface{}{
		"pkg":           cmd.pkg,
		"name":          cmd.name,
		"isPostgresSQL": isPostgresSQL,
		"isRedis":       isRedis,
	}
}

func (cmd *cmdBuilder) Generate() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}

	// project name
	if _, err := os.Stat(cmd.project); os.IsNotExist(err) {
		os.Mkdir(cmd.project, os.ModeDir)
	}

	pathCMD := fmt.Sprintf("%s/%s", cmd.project, cmd.pkg)
	if _, err := os.Stat(pathCMD); os.IsNotExist(err) {
		os.Mkdir(pathCMD, os.ModeDir)
	}

	pathTemplate := fmt.Sprintf("%s/builder/cmd/templates/root.tmpl", pwd)
	t, err := template.ParseFiles(pathTemplate)
	if err != nil {
		log.Print(err)
		os.RemoveAll(cmd.project)
		return
	}

	var buf bytes.Buffer
	data := cmd.generateData()
	if err := t.Execute(&buf, data); err != nil {
		os.RemoveAll(cmd.project)
		return
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		os.RemoveAll(cmd.project)
		return
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/%s/%s.go", pathCMD, cmd.pkg, cmd.name), p, 0644)
	if err != nil {
		os.RemoveAll(cmd.project)
		return
	}
}
