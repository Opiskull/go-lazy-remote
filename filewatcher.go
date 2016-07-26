package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func startFileWatcher(file string) *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("evt:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()
	err = watcher.Add(file)
	if err != nil {
		log.Fatal(err)
	}

	return watcher
}

func initConfigWatcher() func() {
	watcher := startFileWatcher("config.json")
	return func() {
		watcher.Close()
	}
}
