package renderer

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, _ echo.Context) error {
	renderableContent := struct {
		Type    string
		Content any
	}{
		Type:    name, // It is needed to dynamically include templates
		Content: data,
	}

	err := t.templates.ExecuteTemplate(w, "admin_panel_base.html", renderableContent)
	if err != nil {
		return err
	}

	err = t.templates.ExecuteTemplate(w, name, renderableContent)
	if err != nil {
		return err
	}

	return nil
}

func NewTemplates() *Templates {
	t := template.Must(template.ParseGlob("views/*.html"))
	return &Templates{
		templates: t,
	}
}
