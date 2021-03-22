package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/rostislavjadavan/mdwiki/src/storage"
	"github.com/rostislavjadavan/mdwiki/src/ui"
	"net/http"
)

func PageHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("page /" + c.Param("page") + " requested")

		page, err := s.LoadPage(storage.UriToPage(c.Param("page")))
		if err != nil {
			return notFoundPage(err, e, c)
		}

		tpl, err := ui.Render(ui.TemplatePage, page)
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}

func ListHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("page list")

		list, err := s.ListPages()
		if err != nil {
			return errorPage(err, e, c)
		}

		tpl, err := ui.Render(ui.TemplateList, map[string]interface{}{
			"Pages": list,
		})
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}

func CreateHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" {
			name := c.FormValue("name")
			if !storage.ValidateFilename(name) {
				tpl, err := ui.Render(ui.TemplateCreate, map[string]interface{}{
					"Name":       name,
					"Validation": InvalidFilenameValidation,
				})
				if err != nil {
					return errorPage(err, e, c)
				}
				return c.HTML(http.StatusOK, tpl)
			}

			e.Logger.Debug("creating new page " + name)
			page, err := s.CreateNewPage(name)
			if err != nil {
				return errorPage(err, e, c)
			}
			return c.Redirect(http.StatusFound, "edit/"+page.Filename)
		}

		tpl, err := ui.Render(ui.TemplateCreate, map[string]interface{}{
			"Name":       "",
			"Validation": "",
		})
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}

func EditHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("edit /" + c.Param("page"))

		page, err := s.LoadPage(storage.UriToPage(c.Param("page")))
		if err != nil {
			return notFoundPage(err, e, c)
		}

		if c.Request().Method == "POST" {
			c.Logger().Debug("content update of page " + page.Filename)
			content := c.FormValue("content")
			err := s.UpdatePageContent(content, page)
			if err != nil {
				return errorPage(err, e, c)
			}
			return c.Redirect(http.StatusFound, "/"+page.Filename)
		}

		tpl, err := ui.Render(ui.TemplateEdit, page)
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}

func DeleteHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("delete /" + c.Param("page"))

		page, err := s.LoadPage(storage.UriToPage(c.Param("page")))
		if err != nil {
			return notFoundPage(err, e, c)
		}

		tpl, err := ui.Render(ui.TemplateDelete, page)
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}

func DoDeleteHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		page, err := s.LoadPage(storage.UriToPage(c.Param("page")))
		if err != nil {
			return notFoundPage(err, e, c)
		}
		err = s.DeletePage(page)
		if err != nil {
			return errorPage(err, e, c)
		}
		return c.Redirect(http.StatusFound, "/list")
	}
}
