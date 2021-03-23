package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rostislavjadavan/mdwiki/src/ui"
	"net/http"
)

func errorPage(err error, e *echo.Echo, c echo.Context) error {
	e.Logger.Error(err)
	tpl, err := ui.Render(ui.TemplateError, map[string]interface{}{
		"Message": err.Error(),
	})
	if err != nil {
		e.Logger.Error(err)
	}
	return c.HTML(http.StatusInternalServerError, tpl)
}

func notFoundPage(err error, e *echo.Echo, c echo.Context) error {
	e.Logger.Warn(err)
	tpl, err := ui.Render(ui.TemplateNotFound, map[string]interface{}{
		"Message": err.Error(),
	})
	if err != nil {
		e.Logger.Error(err)
	}
	return c.HTML(http.StatusNotFound, tpl)
}
