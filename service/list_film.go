package service

import (
	"bitbucket.org/matiux/archiviofilm/model"
	// "fmt"
)

type ListFilm struct {
	FilmRepository *model.FilmRepository
}

func NewListFilm() *ListFilm {

	filmRepo, _ := model.NewFilmRepository()

	listFilm := ListFilm{FilmRepository: filmRepo}

	return &listFilm
}

func (lf *ListFilm) Execute() []model.Film {

	films := lf.FilmRepository.All()

	return films
}
