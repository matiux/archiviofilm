package model

type ArchivioFilm struct {
	Films []Film
}

func (a *ArchivioFilm) Add(film Film) []Film {

	a.Films = append(a.Films, film)

	return a.Films
}
