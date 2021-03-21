package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Port    string `yaml:"port" env:"PORT" env-default:"8080"`
	Host    string `yaml:"host" env:"HOST" env-default:"localhost"`
	Storage string `yaml:"storage" env:"STORAGE" env-default:".storage"`
}

var config AppConfig

func main() {
	err := cleanenv.ReadConfig("config.yml", &config)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.GET("/static/style.css", StaticHandler(css_style, "text/css", e))
	e.GET("/static/codejar.js", StaticHandler(js_codejar, "text/javascript", e))
	e.GET("/list", ListHandler(e))
	e.GET("/create", CreateHandler(e))
	e.POST("/create", CreateHandler(e))
	e.GET("/edit/:page", EditHandler(e))
	e.POST("/edit/:page", EditHandler(e))
	e.GET("/delete/:page", DeleteHandler(e))
	e.GET("/dodelete/:page", DoDeleteHandler(e))
	e.GET("/", PageHandler(e))
	e.GET("/:page", PageHandler(e))

	e.Logger.Fatal(e.Start(config.Host + ":" + config.Port))
}
