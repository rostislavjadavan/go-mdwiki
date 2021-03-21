package main

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed templates/page.html
var tpl_page string

//go:embed templates/list.html
var tpl_list string

//go:embed templates/create.html
var tpl_create string

//go:embed templates/edit.html
var tpl_edit string

//go:embed templates/delete.html
var tpl_delete string

//go:embed templates/error.html
var tpl_error string

//go:embed templates/not_found.html
var tpl_not_found string

//go:embed templates/head.html
var tpl_head string

//go:embed templates/footer.html
var tpl_footer string

//go:embed templates/style.css
var css_style string

type TemplateData struct {
	Data   interface{}
	Head   string
	Footer string
}

func Render(content string, data interface{}) (string, error) {
	head, err := renderTemplate(tpl_head, nil)
	if err != nil {
		return "", err
	}
	footer, err := renderTemplate(tpl_footer, nil)
	if err != nil {
		return "", err
	}
	tpl, err := renderTemplate(content, TemplateData{
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
