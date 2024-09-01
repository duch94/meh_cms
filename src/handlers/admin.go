package handlers

import "github.com/labstack/echo/v4"

func AdminPanelView(c echo.Context) error {
	err := c.Render(200, "admin_panel", nil)
	if err != nil {
		return err
	}
	return nil
}

func AdminHomeView(c echo.Context) error {
	err := c.Render(200, "admin_home", nil)
	if err != nil {
		return err
	}
	return nil
}
