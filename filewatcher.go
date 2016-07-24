package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func startConfigWatcher() *fsnotify.Watcher {
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
	err = watcher.Add("config.json")
	if err != nil {
		log.Fatal(err)
	}

	return watcher
}
