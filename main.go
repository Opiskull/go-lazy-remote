package main

import (
	"flag"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func main() {
	conf := LoadConfiguration()
	service := NewCommandService(conf.Commands)

	static := web.New()
	static.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/static/", static)

	goji.Use(JSONMiddleware)
	service.Init()

	flag.Set("bind", conf.Listen)
	goji.Serve()
}
