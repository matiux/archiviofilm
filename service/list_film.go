package service

import (
	"bitbucket.org/matiux/archiviofilm/model"
	"fmt"
	"net/url"
)

type ListFilm struct {
	FilmRepository *model.FilmRepository
}

func NewListFilm() *ListFilm {

	filmRepo, _ := model.NewFilmRepository()

	listFilm := ListFilm{FilmRepository: filmRepo}

	return &listFilm
}

func (lf *ListFilm) Execute(filters url.Values) []model.Film {

	fmt.Printf("list_film.go - Filters: %v\nType: %T\n\n", filters, filters)

	films := lf.FilmRepository.All()

	return films
}
