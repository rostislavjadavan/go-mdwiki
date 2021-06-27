package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rostislavjadavan/mdwiki/src/config"
	"github.com/rostislavjadavan/mdwiki/src/storage"
	"github.com/rostislavjadavan/mdwiki/src/webserver"
)

func main() {
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		panic(err)
	}

	s, err := storage.CreateStorage(cfg)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	webserver.SetRoutes(e, s)
	e.Logger.Fatal(e.Start(cfg.Host + ":" + cfg.Port))
}
