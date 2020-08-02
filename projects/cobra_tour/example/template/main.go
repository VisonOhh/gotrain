package main

import (
	"html/template"
	"os"
	"strings"
)

const templateText = `
Output 0: {{title .Name1}}
Output 1: {{upper .Name2}}
Output 2: {{.Name3 | lower}}
`

func main() {
	funcMap := template.FuncMap{
		"title": strings.Title,
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
	}
	tpl, _ := template.New("template_test").Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "vison",
		"Name2": "jimin",
		"Name3": "XIAOYU",
	}
	_ = tpl.Execute(os.Stdout, data)
}
