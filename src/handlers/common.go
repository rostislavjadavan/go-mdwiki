package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rostislavjadavan/go-mdwiki/src/ui"
	"net/http"
)

var InvalidFilenameValidation string = "Invalid filename, valid examples: wiki_page_1.md, flowers-and-animals.md, page106.md"

func errorPage(err error, e *echo.Echo, c echo.Context) error {
	e.Logger.Error(err)
	tpl, _ := ui.Render(ui.TemplateError, map[string]interface{}{
		"Message": err.Error(),
	})
	return c.HTML(http.StatusInternalServerError, tpl)
}
