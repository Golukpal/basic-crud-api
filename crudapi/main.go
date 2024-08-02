package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *director `json:"director"`
}

//  this *director struct has been assosicated with movie struct and director is type of director struct

type director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(movies)
	}
}

// w.Header().set("Content-Type","application/json") is used to set the header of the response to application/json

// here mux.Vars(r) is used to get the parameters from the request and then we are iterating over the movies and checking
//  if the id of the movie is equal to the id of the parameter then we are deleting the movie from the movies slice

// here index+1 is used for when a movie is deleted then the next movie will take its place

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//here params is used to get the parameters from the request and then we are iterating over the movies and checking
// if the id of the movie is equal to the id of the parameter then we are encoding the movie in the response
// here mux.Vars(r) is used to get the parameters from the request and then we are iterating over the movies and checking
//  if the id of the movie is equal to the id of the parameter then we are encoding the movie in the response

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newMovie movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, newMovie)
	json.NewEncoder(w).Encode(newMovie)
}

// here json.NewDecoder(r.body).Decode(&movie) is used to decode the body of the request and store it in the movie variable
// here append is used to append the movie to the movies slice and then we are encoding the movie in the response
/* here movie.ID = strconv.Itoa(rand.Intn(1000000)) is used to generate a random id for the movie and then we
are appending the movie to the movies slice*/
// here _ is used to ignore the error and & is used to get the address of the variable (movie) and store the value in it
// and then decode the value of the movie in the movie variable and r.body is used to get the body of the request

func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return

		}
	}
}

// in update we used both delete and create to update the movie and then we are encoding the movie in the response
// here we are decoding the body of the request and storing it in the movie variable and then we are updating the id of the movie

func main() {

	r := mux.NewRouter()

	movies = append(movies, movie{ID: "1", Isbn: "448743", Title: "Movie One", Director: &director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, movie{ID: "2", Isbn: "448744", Title: "Movie Two", Director: &director{Firstname: "Steve", Lastname: "Smith"}})
	movies = append(movies, movie{ID: "3", Isbn: "448745", Title: "Movie Three", Director: &director{Firstname: "Jane", Lastname: "Doe"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
