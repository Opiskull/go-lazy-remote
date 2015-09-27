package main

import (
	"flag"

	"github.com/opiskull/go-jsonapi"
	"github.com/zenazn/goji"
)

func main() {
	conf := LoadConfiguration()
	service := NewCommandService(conf.Commands)

	goji.DefaultMux = jsonapi.NewJSONMux()
	goji.Use(jsonapi.StaticFiles(conf.StaticFiles))
	service.Init()
	flag.Set("bind", conf.Listen)
	goji.Serve()
}
