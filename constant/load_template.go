package constant

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// init struct for template
type Template struct {
	templates *template.Template
}

// set render used for call template name eg index.html
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func LoadTemplate() *Template {
	//get path to template folder
	template := &Template{
		templates: template.Must(template.ParseGlob("repository/templates/*")),
	}

	//return template path
	return template
}
