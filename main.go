package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// in new slice we are skipping the movie that we want to delete that is at Index -1  position
			// we are adding all movies before the movie and after the movie that, thus skipping the movie that we want to delete
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

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
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(len(movies) + 1)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	// this defines the server
	r := mux.NewRouter()

	// * is used to declare a pointer to a specific type and to access the value at a memory address of an object, struct, or variable. & is used to get the address (reference) of a variable, object, or struct.
	// * is for declaring a pointer to a specific type and accessing the value at a memory address of an object, struct, or variable. & is used to get the address (reference) of a variable, object, or struct.
	// & is used to get the address (reference) of a variable, object, or struct, which can then be assigned to a pointer variable.
	// json.Encoder converts our struct to JSON and json.NewDecoder converts JSON to our struct
	// struct -> JSON -> json.NewEncoder ; JSON -> struct -> json.NewDecoder

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Catch Me if You Can", Director: &Director{Firstname: "Tahseen", Lastname: "Zaman"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438228", Title: "The Prestige", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	// This is starting server at port 8080
	log.Fatal(http.ListenAndServe(":8080", r))

}
