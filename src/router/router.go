package router

import (
	"github.com/duch94/meh_cms/src/handlers"
	"github.com/duch94/meh_cms/src/renderer"
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.Renderer = renderer.NewTemplates()

	e.File("admin_panel.css", "views/admin_panel.css")

	e.GET("/admin", handlers.AdminPanelView)
	e.GET("/admin_home", handlers.AdminHomeView)
	e.GET("/articles", handlers.ArticlesSubView)
	e.GET("/topics", handlers.TopicsSubView)
	e.GET("/users", handlers.UsersSubView)

}
