package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rostislavjadavan/mdwiki/src/api"
	"github.com/rostislavjadavan/mdwiki/src/config"
	"github.com/rostislavjadavan/mdwiki/src/handlers"
	"github.com/rostislavjadavan/mdwiki/src/storage"
	"github.com/rostislavjadavan/mdwiki/src/ui"
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

	// Static files
	e.GET("/static/style.css", handlers.StaticHandler(ui.CssStyle, handlers.MimeCss, e))
	e.GET("/static/codejar.js", handlers.StaticHandler(ui.JavascriptCodeJar, handlers.MimeJavascript, e))
	e.GET("/static/script.js", handlers.StaticHandler(ui.JavascriptScript, handlers.MimeJavascript, e))
	e.GET("/static/favicon.png", handlers.StaticHandler(ui.ImageFaviconPng, handlers.MimePng, e))

	// RPC like API
	e.POST("/api/page.create", api.PageCreateHandler(e, s))
	e.POST("/api/page.update/:page", api.PageUpdateHandler(e, s))
	e.POST("/api/page.delete", api.PageDeleteHandler(e, s))
	e.POST("/api/trash.empty", api.TrashEmptyHandler(e, s))
	e.POST("/api/trash.restore", api.TrashRestoreHandler(e, s))
	e.POST("/api/version.restore", api.VersionRestoreHandler(e, s))

	// UI
	e.GET("/search", handlers.SearchHandler(e, s))
	e.GET("/trash", handlers.TrashHandler(e, s))
	e.GET("/trash/:page", handlers.TrashPageHandler(e, s))
	e.GET("/list", handlers.ListHandler(e, s))
	e.GET("/create", handlers.CreateHandler(e, s))
	e.GET("/edit/:page", handlers.EditHandler(e, s))
	e.GET("/", handlers.PageHandler(e, s))
	e.GET("/:page/version", handlers.PageVersionsHandler(e, s))
	e.GET("/:page/version/:ver", handlers.PageVersionHandler(e, s))
	e.GET("/:page", handlers.PageHandler(e, s))

	e.Logger.Fatal(e.Start(cfg.Host + ":" + cfg.Port))
}
