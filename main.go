package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	conf := LoadConfiguration()

	watcher := startConfigWatcher()
	defer watcher.Close()

	e := echo.New()
	g := e.Group("/api")
	NewCommandService(conf.Commands, g)
	e.Static("/", conf.StaticFiles)
	log.Println("Listening on:" + conf.Listen)
	e.Run(standard.New(conf.Listen))
}
