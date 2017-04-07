package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"bitbucket.org/matiux/archiviofilm/service"
	"github.com/radovskyb/watcher"
)

func main() {

	// usr, err := user.Current()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	const Path = "/mnt/storaMioArchivio/MieiVideo/Film/"
	//const Path = "/home/" + usr.Username + "/Desktop/GoSyncTest"
	const Debug = false

	w := watcher.New()

	// SetMaxEvents to 1 to allow at most 1 Event to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	//w.SetMaxEvents(100)

	var wg sync.WaitGroup
	wg.Add(1)

	CreateFilm := service.NewCreateFilm()
	DeleteFilm := service.NewDeleteFilm()
	RenameFilm := service.NewRenameFilm()

	go func() {

		defer wg.Done()

		for {
			select {
			case event := <-w.Event:
				// Print the event's info.
				//fmt.Println(event)

				if !event.IsDir() {

					switch event.Op {

					case watcher.Write:
						fmt.Println("Event: WRITE [" + event.Path + "]")
						if !Debug {
							CreateFilm.CreateOrUpdateFilm(event.Path)
						}
					case watcher.Move:
						fmt.Println("Event: MOVE")
						splitPath := strings.Split(event.Path, " -> ")
						oldPath := splitPath[0]
						newPath := splitPath[1]
						fmt.Printf("%v -> %v\n", oldPath, newPath)
						if !Debug {
							RenameFilm.Execute(oldPath, newPath)
						}
					case watcher.Create:
						fmt.Println("Event: CREATE [" + event.Path + "]")
						if !Debug {
							CreateFilm.CreateOrUpdateFilm(event.Path)
						}
					case watcher.Remove:
						fmt.Println("Event: REMOVE [" + event.Path + "]")
						if !Debug {
							DeleteFilm.DeleteFilmIfExists(event.Path)
						}
					case watcher.Rename:
						fmt.Println("Event: RENAME")
						splitPath := strings.Split(event.Path, " -> ")
						oldPath := splitPath[0]
						newPath := splitPath[1]
						fmt.Printf("%v -> %v\n", oldPath, newPath)
						if !Debug {
							RenameFilm.Execute(oldPath, newPath)
						}
						// if film := DeleteFilm.DeleteFilmIfExists(oldPath); film != nil {

						// 	CreateFilm.FromExist(film, newPath)
						// }
					case watcher.Chmod:
						fmt.Println("Event: CHMODED [" + event.Path + "]")
					}
				}

			case err := <-w.Error:
				log.Fatalln(err)
			}
		}
	}()

	// Watch test_folder recursively for changes.
	if err := w.AddRecursive(Path); err != nil {
		log.Fatalln(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}

	wg.Wait()
}
