package main

import (
	"fmt"
	"os"
	"path/filepath"

	"bitbucket.org/matiux/archiviofilm/service"
	"github.com/fatih/color"
)

func main() {

	yellow := color.New(color.FgYellow).Add(color.Underline)

	searchDir := "/mnt/storaMioArchivio/MieiVideo/Film/"

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
