package libos

import (
	"bytes"
	"fmt"
	"html/template"
)

func ExecuteTemplate(templateBytes []byte, data interface{}) ([]byte, error) {
	t, err := template.New("template").
		Funcs(template.FuncMap{
			"RawHTML": func(s string) template.HTML {
				return template.HTML(s)
			},
		}).
		Parse(string(templateBytes))
	if err != nil {
		return nil, fmt.Errorf("error parsing template: %w", err)
	}

	var buff bytes.Buffer
	err = t.Execute(&buff, data)
	if err != nil {
		return nil, fmt.Errorf("error executing template: %w", err)
	}

	return buff.Bytes(), nil
}
