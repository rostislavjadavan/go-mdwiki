package api

import (
	"github.com/labstack/echo/v4"
	"github.com/rostislavjadavan/mdwiki/src/storage"
	"net/http"
)

func EmptyTrashHandler(e *echo.Echo, s *storage.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		e.Logger.Debug("api empty trash")
		err := s.EmptyTrash()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		}
		return c.JSON(http.StatusOK, RedirectResponse{Redirect: "/trash"})
	}
}
