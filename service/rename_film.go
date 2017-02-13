package service

import (
	"bitbucket.org/matiux/archiviofilm/model"
	"github.com/fatih/color"
	"time"
)

type RenameFilm struct {
	FilmRepository *model.FilmRepository
}

func NewRenameFilm() *RenameFilm {

	filmRepo, _ := model.NewFilmRepository()

	renameFilm := RenameFilm{FilmRepository: filmRepo}

	return &renameFilm
}

func (r *RenameFilm) Execute(oldName string, newName string) {

	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)

	if film, exists := r.FilmRepository.FindFilmByName(oldName); exists {

		film.File = newName
		film.UpdatedAt = time.Now()

		r.FilmRepository.Save(film)

		green.Printf("Film aggiornato: %v -> %v\n", oldName, newName)

	} else {

		red.Printf("Film non esistente: %v\n", oldName)
	}
}
