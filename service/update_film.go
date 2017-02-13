package service

import (
	"bitbucket.org/matiux/archiviofilm/model"
	"github.com/fatih/color"
	"time"
)

type UpdateFilm struct {
	FilmRepository *model.FilmRepository
}

func NewUpdateFilm() *UpdateFilm {

	filmRepo, _ := model.NewFilmRepository()

	updateFilm := UpdateFilm{FilmRepository: filmRepo}

	return &updateFilm
}

func (u *UpdateFilm) updateSeen(filmId string, seen bool) {

	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)

	if film, exists := u.FilmRepository.FindFilmById(filmId); exists {

		film.Seen = seen
		film.UpdatedAt = time.Now()

		u.FilmRepository.Save(film)

		green.Printf("Film aggiornato [Seen]\n")

	} else {

		red.Printf("Film non esistente: %v\n", filmId)
	}
}
