package handlers

import "github.com/labstack/echo/v4"

func TopicsSubView(c echo.Context) error {
	err := c.Render(200, "topics", nil)
	if err != nil {
		return err
	}
	return nil
}
