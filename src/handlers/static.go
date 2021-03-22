package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func StaticHandler(content string, contentType string, e *echo.Echo) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.Response().Header().Add("Content-Type", contentType)
		return c.String(http.StatusOK, content)
	}
}
