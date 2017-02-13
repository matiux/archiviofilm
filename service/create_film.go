package service

import (
	"bitbucket.org/matiux/archiviofilm/model"
	"github.com/fatih/color"
)

type CreateFilm struct {
	FilmRepository *model.FilmRepository
}

func NewCreateFilm() *CreateFilm {

	filmRepo, _ := model.NewFilmRepository()

	createFilm := CreateFilm{FilmRepository: filmRepo}

	return &createFilm
}

func (c *CreateFilm) CreateOrUpdateFilm(path string) {

	film, exists := c.FilmRepository.FindFilmByName(path)

	green := color.New(color.FgGreen)
	// yellow := color.New(color.FgYellow).Add(color.Underline)
	red := color.New(color.FgRed)

	if !exists {

		film := model.NewFilm(path, false, false)

		c.FilmRepository.CreateFilm(film)

		green.Println(" - Record creato: ", path)

	} else {

		film.File = path

		c.FilmRepository.Save(film)

		red.Println(" - Record aggiornato: ", path)
	}
}

func (c *CreateFilm) FromExist(newFile string, oldFile *model.Film) {

	green := color.New(color.FgGreen)
	// yellow := color.New(color.FgYellow).Add(color.Underline)
	//red := color.New(color.FgRed)

	_, exists := c.FilmRepository.FindFilmByName(oldFile.File)

	if !exists {

		var withChecksum bool

		if 0 < len(oldFile.Checksum) {

			withChecksum = true
		} else {

			withChecksum = false
		}

		film := model.NewFilm(newFile, oldFile.Seen, withChecksum)

		c.FilmRepository.CreateFilm(film)

		green.Println(" - Record creato: ", newFile)
	}
}
