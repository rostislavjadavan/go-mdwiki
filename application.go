package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rostislavjadavan/go-mdwiki/src/config"
	"github.com/rostislavjadavan/go-mdwiki/src/handlers"
	"github.com/rostislavjadavan/go-mdwiki/src/storage"
	"github.com/rostislavjadavan/go-mdwiki/src/ui"
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
	e.GET("/static/style.css", handlers.StaticHandler(ui.CssStyle, handlers.MimeCss, e))
	e.GET("/static/codejar.js", handlers.StaticHandler(ui.JavascriptCodeJar, handlers.MimeJavascript, e))
	e.GET("/list", handlers.ListHandler(e, s))
	e.GET("/create", handlers.CreateHandler(e, s))
	e.POST("/create", handlers.CreateHandler(e, s))
	e.GET("/edit/:page", handlers.EditHandler(e, s))
	e.POST("/edit/:page", handlers.EditHandler(e, s))
	e.GET("/delete/:page", handlers.DeleteHandler(e, s))
	e.GET("/dodelete/:page", handlers.DoDeleteHandler(e, s))
	e.GET("/", handlers.PageHandler(e, s))
	e.GET("/:page", handlers.PageHandler(e, s))

	e.Logger.Fatal(e.Start(cfg.Host + ":" + cfg.Port))
}
