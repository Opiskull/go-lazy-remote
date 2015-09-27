package main

import (
	"flag"

	"github.com/zenazn/goji"
)

func main() {
	conf := LoadConfiguration()

	service := NewCommandService(conf.Commands)
	goji.Handle("/api/*", service.Mux)

	files := NewFileService(conf.StaticFiles)
	goji.Handle("/*", files.Mux)

	flag.Set("bind", conf.Listen)
	goji.Serve()
}
