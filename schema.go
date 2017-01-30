package main

import (
	"bitbucket.org/matiux/archiviofilm/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db, err := gorm.Open("mysql", "root:281285@/archivio_film?charset=utf8&parseTime=True&loc=Local")

	if err != nil {

		fmt.Println("Errore DB" + err.Error())
	}

	db.AutoMigrate(&model.Film{})

	db.Model(&model.Film{}).AddUniqueIndex("idx_film_name", "File")
}
