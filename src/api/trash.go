package api

import (
	"github.com/labstack/echo/v4"
	"github.com/rostislavjadavan/mdwiki/src/storage"
	"net/http"
)

func TrashEmptyHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("api empty trash")
		err := s.TrashEmpty()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, RedirectResponse{Redirect: "/trash"})
	}
}

func TrashRestoreHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(FilenameRequest)
		err := c.Bind(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		}
		req.Filename = storage.FixPageExtension(req.Filename)

		e.Logger.Debugf("api trash restore '%s'", req.Filename)

		page, err := s.TrashPage(storage.FixPageExtension(req.Filename))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}

		err = s.TrashRestore(page)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}

		return c.JSON(http.StatusOK, RedirectResponse{Redirect: "/" + page.Filename})
	}
}
