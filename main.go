package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	conf := LoadConfiguration()
	watcher := startConfigWatcher()
	defer watcher.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(conf.Cors))
	e.Use(middleware.Secure())

	g := e.Group("/api")
	NewCommandService(conf.CommandFilesFolder, g)
	e.Static("/", conf.StaticFilesFolder)
	log.Println("Listening on:" + conf.Listen)
	e.Run(standard.New(conf.Listen))
}
