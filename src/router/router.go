package router

import (
	"github.com/duch94/meh_cms/src/handlers"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	return err
}

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func InitRouter(e *echo.Echo) {
	e.Renderer = newTemplates()

	e.File("admin_panel.css", "views/admin_panel.css")

	e.GET("/admin", handlers.AdminPanelView)

	e.GET("/admin_home", handlers.AdminHomeView)
	e.GET("/articles", handlers.ArticlesSubView)
	e.GET("/topics", handlers.TopicsSubView)
	e.GET("/users", handlers.UsersSubView)

}
