package cmd

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

//go:embed templates/go_workflow_builder.go.tmpl
var goWorkflowTemplate string

type TemplateWorkflowGeneratorHelper struct {
	Templates map[string]string
}

func (t *TemplateWorkflowGeneratorHelper) Generate(info GeneratedInfo) (map[string]string, error) {
	files := map[string]string{}
	if t.Templates == nil {
		return files, nil
	}

	for file, t := range t.Templates {
		content, err := genFromTemplate(file, t, info)
		if err != nil {
			return nil, err
		}

		// can use a template, but it's simple for now
		fileName, err := genFromTemplate("file name for "+file, file, info)
		if err != nil {
			return nil, err
		}
		files[fileName] = content
	}

	return files, nil
}

func genFromTemplate(name, rawTemplate string, info GeneratedInfo) (string, error) {
	t, err := template.New(name).Funcs(template.FuncMap{
		"LowerFirst": func(s string) string {
			if len(s) == 0 {
				return s
			}
			return strings.ToLower(s[:1]) + s[1:]
		},
		"ToSnake": strcase.ToSnake,
	}).Parse(rawTemplate)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	err = t.Execute(buf, info)
	return buf.String(), err
}
