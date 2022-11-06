package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	Id      string
	Isbn    string
	Title   string
	Creator *Creator
}
type Creator struct {
	Firstname string
	Lastname  string
}

var movies []Movie

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}
func getById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range movies {
		if item.Id == param["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}

}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range movies {
		if param["id"] == item.Id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range movies {
		var movie Movie
		if item.Id == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = param["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{Id: "1", Isbn: "212412", Title: "Titanic", Creator: &Creator{Firstname: "James", Lastname: "Cameron"}})
	movies = append(movies, Movie{Id: "2", Isbn: "152474", Title: "the Green Mile", Creator: &Creator{Firstname: "Steven", Lastname: "King"}})
	movies = append(movies, Movie{Id: "3", Isbn: "918643", Title: "The mask", Creator: &Creator{Firstname: "Chuck", Lastname: "Rassel"}})
	r.HandleFunc("/movies", getAll).Methods("GET")
	r.HandleFunc("/movies/{id}", getById).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	fmt.Println("Server is listening")
	log.Fatal(http.ListenAndServe(":8080", r))

}
