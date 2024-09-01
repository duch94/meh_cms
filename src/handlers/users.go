package handlers

import "github.com/labstack/echo/v4"

func UsersSubView(c echo.Context) error {
	err := c.Render(200, "users", nil)
	if err != nil {
		return err
	}
	return nil
}
