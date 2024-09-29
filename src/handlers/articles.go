package handlers

import (
	"github.com/duch94/meh_cms/src/repo"
	"github.com/labstack/echo/v4"
)

func ArticlesSubView(c echo.Context) error {
	content, err := repo.GetArticles()
	if err != nil {
		return err
	}

	err = c.Render(200, "articles", content)
	if err != nil {
		return err
	}

	return nil
}
