package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rostislavjadavan/mdwiki/src/storage"
	"github.com/rostislavjadavan/mdwiki/src/ui"
	"net/http"
)

func PageVersionsHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("versions")

		pageUri := c.Param("page")
		if pageUri != storage.FixPageExtension(pageUri) {
			return c.Redirect(http.StatusPermanentRedirect, "/"+storage.FixPageExtension(pageUri)+"/version")
		}

		page, err := s.Page(pageUri)
		if err != nil {
			return notFoundPage(err, e, c)
		}

		list, err := s.VersionsList(page.Filename)
		if err != nil {
			return errorPage(err, e, c)
		}

		tpl, err := ui.Render(ui.TemplateVersions, map[string]interface{}{
			"Page":  page,
			"Pages": list,
		})
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}

func PageVersionHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("version page /" + c.Param("ver") + " requested")

		pageUri := c.Param("ver")
		page, err := s.VersionPage(pageUri)
		if err != nil {
			return notFoundPage(err, e, c)
		}

		tpl, err := ui.Render(ui.TemplatePageVersion, page)
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}
