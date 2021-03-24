package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rostislavjadavan/mdwiki/src/storage"
	"github.com/rostislavjadavan/mdwiki/src/ui"
	"net/http"
)

func TrashHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("trash")

		list, err := s.TrashList()
		if err != nil {
			return errorPage(err, e, c)
		}

		tpl, err := ui.Render(ui.TemplateTrash, map[string]interface{}{
			"Pages": list,
		})
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}

func TrashPageHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("trash page /" + c.Param("page") + " requested")

		pageUri := c.Param("page")
		if pageUri != storage.FixPageExtension(pageUri) {
			return c.Redirect(http.StatusPermanentRedirect, "/trash/"+storage.FixPageExtension(pageUri))
		}

		page, err := s.TrashPage(pageUri)
		if err != nil {
			return notFoundPage(err, e, c)
		}

		tpl, err := ui.Render(ui.TemplatePageTrash, page)
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}
