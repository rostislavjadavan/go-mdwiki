package api

import (
	"github.com/labstack/echo/v4"
	"github.com/rostislavjadavan/mdwiki/src/storage"
	"net/http"
)

func VersionRestoreHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(FilenameRequest)
		err := c.Bind(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		}
		e.Logger.Debugf("api version restore '%s'", req.Filename)

		page, err := s.VersionPage(req.Filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}

		err = s.VersionRestore(page)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}

		return c.JSON(http.StatusOK, RedirectResponse{Redirect: "/" + page.Name})
	}
}
