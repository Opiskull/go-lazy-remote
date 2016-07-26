package main

import "github.com/spf13/viper"

const (
	ConfStaticFolder   string = "staticFolder"
	ConfCommandsFolder string = "commandsFolder"
	ConfListen         string = "listen"
	ConfLogFilename    string = "log.filename"
	ConfCorsConfig     string = "cors"
)

func initConfig() {
	viper.SetConfigFile("./config.json")
	viper.SetDefault("listen", ":8000")
	viper.SetDefault("staticFolder", "./static")
	viper.SetDefault("commandsFolder", "./commands")
	viper.SetDefault("log.filename", "go-lazy-remote.log")
	viper.ReadInConfig()
}
