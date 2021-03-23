package ui

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed templates/page.html
var TemplatePage string

//go:embed templates/list.html
var TemplateList string

//go:embed templates/create.html
var TemplateCreate string

//go:embed templates/edit.html
var TemplateEdit string

//go:embed templates/delete.html
var TemplateDelete string

//go:embed templates/error.html
var TemplateError string

//go:embed templates/not_found.html
var TemplateNotFound string

//go:embed templates/head.html
var TemplateHtmlHead string

//go:embed templates/footer.html
var TemplateFooter string

//go:embed assets/style.css
var CssStyle string

//go:embed assets/codejar.js
var JavascriptCodeJar string

//go:embed assets/script.js
var JavascriptScript string

//go:embed assets/favicon.png
var ImageFaviconPng string

type templateData struct {
	Data   interface{}
	Head   string
	Footer string
}

func Render(content string, data interface{}) (string, error) {
	head, err := renderTemplate(TemplateHtmlHead, nil)
	if err != nil {
		return "", err
	}
	footer, err := renderTemplate(TemplateFooter, nil)
	if err != nil {
		return "", err
	}
	tpl, err := renderTemplate(content, templateData{
		Head:   head,
		Data:   data,
		Footer: footer,
	})
	if err != nil {
		return "", err
	}
	return tpl, nil
}

func renderTemplate(content string, data interface{}) (string, error) {
	tc, err := template.New("tpl").Parse(content)
	if err != nil {
		return "", err
	}
	tpl := new(bytes.Buffer)
	err = tc.Execute(tpl, data)
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
