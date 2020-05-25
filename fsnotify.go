package main

import (
	"log"
	"mt/vcs"
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

						// git add
						git := vcs.ByCmd("git")
						if err := git.AddAll(dataDir()); err != nil {
							panic(err)
						}

						// git commit
						if err := git.CommitAll(dataDir()); err != nil {
							panic(err)
						}

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
