package main

import (
	"bitbucket.org/matiux/archiviofilm/service"
	"fmt"
	"github.com/radovskyb/watcher"
	"log"
	"strings"
	"sync"
	"time"
)

func main() {

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

	go func() {

		defer wg.Done()

		for {
			select {
			case event := <-w.Event:
				// Print the event's info.
				//fmt.Println(event)

				// fileType := func() string {
				// 	if event.IsDir() {
				// 		return "Directory"
				// 	} else {
				// 		return "File"
				// 	}
				// }

				// fmt.Println(fileType())

				// Print out the file name with a message based on the event type.

				if !event.IsDir() {

					switch event.Op {

					case watcher.Write:
						fmt.Printf("Event: WRITE")
						CreateFilm.CreateOrUpdateFilm(event.Path)
					case watcher.Create:
						fmt.Printf("Event: CREATE")
						CreateFilm.CreateOrUpdateFilm(event.Path)
					case watcher.Remove:
						fmt.Printf("Event: REMOVE")
						DeleteFilm.DeleteFilmIfExists(event.Path)
					case watcher.Rename:
						fmt.Printf("Event: RENAME")
						splitPath := strings.Split(event.Path, " -> ")
						oldPath := splitPath[0]
						newPath := splitPath[1]

						if film := DeleteFilm.DeleteFilmIfExists(oldPath); film != nil {

							CreateFilm.CopyFilm(newPath, film)
						}

					case watcher.Chmod:
						fmt.Printf("Event: CHMODED")
						// s := fmt.Sprintf("Chmoded %v: %v", fileType(), event.Path)
						// fmt.Println(s)
					}
				}

			case err := <-w.Error:
				log.Fatalln(err)
			}
		}
	}()

	// Watch test_folder recursively for changes.
	if err := w.Add("/home/iux/Desktop/GoSyncTest"); err != nil {
		log.Fatalln(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}

	wg.Wait()
}
