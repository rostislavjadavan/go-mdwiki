package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var InvalidFilenameValidation string = "Invalid filename, valid examples: wiki_page_1.md, flowers-and-animals.md, page106.md"

func errorPage(err error, e *echo.Echo, c echo.Context) error {
	e.Logger.Error(err)
	tpl, _ := Render(tpl_error, map[string]interface{}{
		"Message": err.Error(),
	})
	return c.HTML(http.StatusInternalServerError, tpl)
}

func PageHandler(e *echo.Echo) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("page /" + c.Param("page") + " requested")

		page, err := LoadPage(UriToPage(c.Param("page")))
		if err != nil {
			e.Logger.Warn(err)
			return c.HTML(http.StatusNotFound, tpl_not_found)
		}

		tpl, err := Render(tpl_page, page)
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}

func ListHandler(e *echo.Echo) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("page list")

		list, err := ListPages()
		if err != nil {
			return errorPage(err, e, c)
		}

		tpl, err := Render(tpl_list, map[string]interface{}{
			"Pages": list,
		})
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}

func CreateHandler(e *echo.Echo) func(c echo.Context) error {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" {
			name := c.FormValue("name")
			if !ValidateFilename(name) {
				tpl, err := Render(tpl_create, map[string]interface{}{
					"Name":       name,
					"Validation": InvalidFilenameValidation,
				})
				if err != nil {
					return errorPage(err, e, c)
				}
				return c.HTML(http.StatusOK, tpl)
			}

			e.Logger.Debug("creating new page " + name)
			page, err := CreateNewPage(name)
			if err != nil {
				return errorPage(err, e, c)
			}
			return c.Redirect(http.StatusFound, "edit/"+page.Filename)
		}

		tpl, err := Render(tpl_create, map[string]interface{}{
			"Name":       "",
			"Validation": "",
		})
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}

func EditHandler(e *echo.Echo) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("edit /" + c.Param("page"))

		page, err := LoadPage(UriToPage(c.Param("page")))
		if err != nil {
			e.Logger.Warn(err)
			return c.HTML(http.StatusNotFound, tpl_not_found)
		}

		if c.Request().Method == "POST" {
			c.Logger().Debug("content update of page " + page.Filename)
			content := c.FormValue("content")
			err := page.UpdateContent(content)
			if err != nil {
				return errorPage(err, e, c)
			}
			return c.Redirect(http.StatusFound, "/"+page.Filename)
		}

		tpl, err := Render(tpl_edit, page)
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}

func DeleteHandler(e *echo.Echo) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("delete /" + c.Param("page"))

		page, err := LoadPage(UriToPage(c.Param("page")))
		if err != nil {
			e.Logger.Warn(err)
			return c.HTML(http.StatusNotFound, tpl_not_found)
		}

		tpl, err := Render(tpl_delete, page)
		if err != nil {
			return errorPage(err, e, c)
		}

		return c.HTML(http.StatusOK, tpl)
	}
}

func DoDeleteHandler(e *echo.Echo) func(c echo.Context) error {
	return func(c echo.Context) error {
		page, err := LoadPage(UriToPage(c.Param("page")))
		if err != nil {
			e.Logger.Warn(err)
			return c.HTML(http.StatusNotFound, tpl_not_found)
		}
		err = page.Delete()
		if err != nil {
			return errorPage(err, e, c)
		}
		return c.Redirect(http.StatusFound, "/list")
	}
}

func StaticHandler(content string, contentType string, e *echo.Echo) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.Response().Header().Add("Content-Type", contentType)
		return c.String(http.StatusOK, content)
	}
}
