package service

import (
	"bitbucket.org/matiux/archiviofilm/model"
	"github.com/fatih/color"
)

type DeleteFilm struct {
	FilmRepository *model.FilmRepository
}

func NewDeleteFilm() *DeleteFilm {

	filmRepo, _ := model.NewFilmRepository()

	deleteFilm := DeleteFilm{FilmRepository: filmRepo}

	return &deleteFilm
}

func (d *DeleteFilm) DeleteFilmIfExists(file string) *model.Film {

	film, exists := d.FilmRepository.FindFilmByName(file)

	green := color.New(color.FgGreen)
	// yellow := color.New(color.FgYellow).Add(color.Underline)
	red := color.New(color.FgRed)

	if exists {

		d.FilmRepository.DeleteFilm(film)

		green.Println(" - Record cancellato: ", file)

		return film

	} else {

		red.Println(" - Record non presente: ", file)

		return nil
	}
}
