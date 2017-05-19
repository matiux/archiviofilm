package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"bitbucket.org/matiux/archiviofilm/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var responseMessage = make(map[string]string)

func filmListEndPoint(w http.ResponseWriter, req *http.Request) {

	filters := req.URL.Query().Get("filters")
	unseen := req.URL.Query().Get("unseen")
	sort := req.URL.Query().Get("sort")

	//fmt.Printf("%v, %T\n", filters, filters)
	//fmt.Printf("%v, %T\n", sort, sort)
	//fmt.Printf("%v, %T\n", unseen, unseen)

	films := service.NewListFilm().Execute(filters, unseen, sort)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(films); err != nil {
		panic(err)
	}

	fmt.Printf("Endpoint Hit: %v [FilmListEndPoint()]\n", html.EscapeString(req.URL.Path))
}

func filmUpdateEndPoint(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	needs := []string{"Seen"}

	if validated, err := service.DecodeAndValidate(req, needs); err != nil {

		responseMessage["status"] = err.Error()

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseMessage)

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
	s.HandleFunc("/film", filmListEndPoint).Methods("GET")
	s.HandleFunc("/film/{film}", filmUpdateEndPoint).Methods("PATCH")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "PATCH"},
	})

	handler := c.Handler(s)

	log.Fatal(http.ListenAndServe(":8080", handler))

	fmt.Println("localhost:8080 is listening")
}
