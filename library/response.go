package library

import "html/template"

func NewTemplate() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}
