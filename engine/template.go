package engine

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/playwright-community/playwright-go"
)

func applyTemplate(text string, vars Vars, page playwright.Page) (string, error) {
	tmpl, err := template.New("template").Parse(text)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %s", err)
	}
	variables := vars.LiveSnapshot()
	variables["page"] = page
	output := bytes.NewBufferString("")
	err = tmpl.Execute(output, variables)
	if err != nil {
		return "", fmt.Errorf("failed to execute template using vars snapshot: %s", err)
	}
	return output.String(), nil
}
