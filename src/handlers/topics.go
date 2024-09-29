package handlers

import (
	"github.com/duch94/meh_cms/src/repo"
	"github.com/labstack/echo/v4"
)

func TopicsSubView(c echo.Context) error {
	data, err := repo.GetTopics()
	if err != nil {
		return err
	}

	err = c.Render(200, "topics", data)
	if err != nil {
		return err
	}
	return nil
}
