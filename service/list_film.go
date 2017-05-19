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

func (lf *ListFilm) Execute(filterLIst string, unseen string, sort string) []model.Film {

	filters := strings.Split(filterLIst, ",")

	var unseenCheck bool

	if "true" == unseen {

		unseenCheck = true

	} else {

		unseenCheck = false
	}

	films := lf.FilmRepository.All(filters, unseenCheck, sort)

	return films
}
