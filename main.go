package main

import (
	"io"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	conf := LoadConfiguration()

	defer initLogging(conf)()
	defer initConfigWatcher()()
	
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

func initLogging(conf *Configuration) func() {
	f, err := os.OpenFile(conf.LogFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.Print("Logging to " + conf.LogFileName)
	multi := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)
	return func() {
		if f != nil {
			f.Close()
		}
	}
}
