package model

import (
	//"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type FilmRepository struct {
	db *gorm.DB
}

func NewFilmRepository() (*FilmRepository, *gorm.DB) {

	db, err := gorm.Open("mysql", "root:281285@/archivio_film?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		fmt.Println("Errore DB" + err.Error())
	}

	filmRepo := FilmRepository{db: db}

	return &filmRepo, db
}

func (filmRepo *FilmRepository) CloseDbConnection() {
	filmRepo.db.Close()
}

func (filmRepo *FilmRepository) FindFilmById(filmId string) (*Film, bool) {

	exists := false

	film := new(Film)

	filmRepo.db.Where(&Film{Id: filmId}).First(&film)

	if film.Id != "" {
		exists = true
	}

	return film, exists
}

func (filmRepo *FilmRepository) FindFilmByName(fileName string) (*Film, bool) {
	//, error
	// if 0 == 0 {
	// 	err := errors.New("Qualcosa non va....")
	// }

	exists := false

	film := new(Film)

	filmRepo.db.Where(&Film{File: fileName}).First(&film)

	if film.Id != "" {
		exists = true
	}

	return film, exists
}

func (filmRepo *FilmRepository) CreateFilm(film Film) Film {

	filmRepo.db.NewRecord(film)

	filmRepo.db.Create(&film)

	return film
}

// func UpdateFilm(film Film, name string) (Film, error) {
// }

func (filmRepo *FilmRepository) All(filters []string, unseenCheck bool, sort string) []Film {

	fmt.Printf("Filters: %v\nSort: %v\nUnseenCheck: %v\n\n", filters, sort, unseenCheck)

	var films []Film

	query := filmRepo.db.Order("File " + sort)

	for _, element := range filters {

		query = query.Where("file LIKE ?", "%"+element+"%")
	}

	if unseenCheck {

		query = query.Where("seen = 0")
	}

	if query.Find(&films).Error != nil {

		//...
	}

	return films
}

func (filmRepo *FilmRepository) DeleteFilm(film *Film) (bool, error) {

	err := filmRepo.db.Delete(&film).Error

	if err != nil {

		fmt.Println("Problemi con la cancellazione del record, err=%s", err)

		return false, err
	}

	return true, err
}

func (filmRepo *FilmRepository) Save(film *Film) {

	filmRepo.db.Save(&film)
}
