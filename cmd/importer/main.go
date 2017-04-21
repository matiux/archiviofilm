package main

import (
	"fmt"
	"os"
	"path/filepath"

	"bitbucket.org/matiux/archiviofilm/service"
	"github.com/fatih/color"
)

func main() {

	var searchDir string

	if 1 < len(os.Args) {

		searchDir = os.Args[1]

		if _, err := os.Stat(searchDir); err != nil {

			if os.IsNotExist(err) {
				fmt.Println("'" + searchDir + "' doesn't exist")
				os.Exit(1)
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	} else {

		searchDir = "/mnt/storaMioArchivio/MieiVideo/Film/"
	}

	yellow := color.New(color.FgYellow).Add(color.Underline)

	CreateFilm := service.NewCreateFilm()

	defer CreateFilm.FilmRepository.CloseDbConnection()

	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {

		if !f.IsDir() {

			yellow.Print("Found " + path)

			if err != nil {

				fmt.Println("Errore " + err.Error())
			}

			CreateFilm.CreateOrUpdateFilm(path)
		}

		return nil
	})
}
