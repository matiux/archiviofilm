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

type RequestFilm struct {
	Seen bool
}

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

	needs := []string{"Seen"}

	service.DecodeAndValidate(req, needs)

	//UpdateFilm := service.NewUpdateFilm()

	//w.Header().Set("Content-Type", "application/json")

	// var requestFilm RequestFilm

	// params := mux.Vars(req)

	//fmt.Println(mux.Vars(req)["film"])

	// if req.Body == nil {
	// 	http.Error(w, "Please send a request body", 400)
	// 	return
	// }

	// err := json.NewDecoder(req.Body).Decode(&requestFilm)

	// if err != nil {
	// 	http.Error(w, "Decode error: "+err.Error(), 400)
	// 	return
	// }
	//fmt.Println(req.Body)
	// fmt.Println(requestFilm.Seen)
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	s := router.PathPrefix("/api/v1").Subrouter()
	s.Host("localhost:8080")
	s.HandleFunc("/films", FilmListEndPoint).Methods("GET")
	s.HandleFunc("/films/{film}", FilmUpdateEndPoint).Methods("PATCH")
	fmt.Println("localhost:8080 is listening")
	log.Fatal(http.ListenAndServe(":8080", router))
}
