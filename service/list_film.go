package service

import (
	"bitbucket.org/matiux/archiviofilm/model"
	//"fmt"
	"strings"
)

type ListFilm struct {
	FilmRepository *model.FilmRepository
}

func NewListFilm() *ListFilm {

	filmRepo, _ := model.NewFilmRepository()

	listFilm := ListFilm{FilmRepository: filmRepo}

	return &listFilm
}

func (lf *ListFilm) Execute(filterLIst string, sort string) []model.Film {

	filters := strings.Split(filterLIst, ",")

	films := lf.FilmRepository.All(filters, sort)

	return films
}
