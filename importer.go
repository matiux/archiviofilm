package main

import (
	"bitbucket.org/matiux/archiviofilm/service"
	"fmt"
	"github.com/fatih/color"
	"os"
	"path/filepath"
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
