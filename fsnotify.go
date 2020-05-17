package main

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// Monitor file change event in a directory
func Monitor(dir string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write ||
					event.Op&fsnotify.Remove == fsnotify.Remove ||
					event.Op&fsnotify.Create == fsnotify.Create ||
					event.Op&fsnotify.Rename == fsnotify.Rename {
					// Vim will emit 'rename' event, rather than 'write'.
					// Further detail: https://github.com/fsnotify/fsnotify/issues/54
					if !strings.Contains(event.Name, ".swp") {
						// Ignore vim .swp file

						// git add && git commit
						AddAndCommit()
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				panic(err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		panic(err)
	}
	<-done
}
