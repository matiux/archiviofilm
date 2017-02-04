package main

import (
	"bitbucket.org/matiux/archiviofilm/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	// "io"
	"log"
	"net/http"
)

func FilmListEndPoint(w http.ResponseWriter, req *http.Request) {

	films := service.NewListFilm().Execute()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(films); err != nil {
		panic(err)
	}

	//serialized, _ := json.Marshal(films)

	// io.WriteString(w, string(res2B))

	//io.WriteString(w, string(serialized))

	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))

	fmt.Printf("Endpoint Hit: %v [FilmListEndPoint()]\n", html.EscapeString(req.URL.Path))
}

func FilmUpdateEndPoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	fmt.Println(params)
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.Host("localhost:8080")
	router.HandleFunc("/films", FilmListEndPoint).Methods("GET")
	router.HandleFunc("/films/{film}", FilmUpdateEndPoint).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8080", router))
}
