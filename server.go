package main

import (
	"encoding/json"
	// "fmt"
	"bitbucket.org/matiux/archiviofilm/service"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func FilmListEndPoint(w http.ResponseWriter, req *http.Request) {

	films := service.NewListFilm().Execute()

	// res2B, _ := json.Marshal(films)

	// io.WriteString(w, string(res2B))
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/films", FilmListEndPoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
