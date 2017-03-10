package main

import (
	"bitbucket.org/matiux/archiviofilm/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

var ResponseMessage = make(map[string]string)

func FilmListEndPoint(w http.ResponseWriter, req *http.Request) {

	filters := req.URL.Query().Get("filters")
	sort := req.URL.Query().Get("sort")

	// fmt.Printf("%v, %T\n", filters, filters)
	// fmt.Printf("%v, %T\n", sort, sort)

	films := service.NewListFilm().Execute(filters, sort)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(films); err != nil {
		panic(err)
	}

	fmt.Printf("Endpoint Hit: %v [FilmListEndPoint()]\n", html.EscapeString(req.URL.Path))
}

func FilmUpdateEndPoint(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	needs := []string{"Seen"}

	if validated, err := service.DecodeAndValidate(req, needs); err != nil {

		ResponseMessage["status"] = err.Error()

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseMessage)

	} else {

		UpdateFilm := service.NewUpdateFilm()
		film := UpdateFilm.UpdateSeen(mux.Vars(req)["film"], validated.Seen)

		if film != nil {

			//ResponseMessage["status"] = "updated"
			json.NewEncoder(w).Encode(film)
		}
	}
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	s := router.PathPrefix("/api/v1").Subrouter()
	s.Host("localhost:8080")
	s.HandleFunc("/film", FilmListEndPoint).Methods("GET")
	s.HandleFunc("/film/{film}", FilmUpdateEndPoint).Methods("PATCH")

	fmt.Println("localhost:8080 is listening")

	log.Fatal(http.ListenAndServe(":8080", router))
}
