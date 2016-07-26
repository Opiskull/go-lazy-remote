package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

func main() {
	initConfig()

	defer initLogging(viper.GetString(ConfLogFilename))()
	defer initConfigWatcher()()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(middleware.CORSWithConfig(viper.Get(ConfCorsConfig).(middleware.CORSConfig)))
	e.Use(middleware.Secure())
	g := e.Group("/api")
	NewCommandService(viper.GetString(ConfCommandsFolder), g)
	e.Static("/", viper.GetString(ConfStaticFolder))
	log.Println("Listening on:" + viper.GetString(ConfListen))
	e.Run(standard.New(viper.GetString(ConfListen)))
}
