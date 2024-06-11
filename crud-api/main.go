package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


type Movie struct {
	ID string `json:"id"`
	Title string `json:"title"`
	ISBN string `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct {
	Name string `json:"name"`
}

var movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	found := false

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1:]...)
			json.NewEncoder(w).Encode(movies)
			found = true
			break
		}
	}

	if !found {
        http.Error(w, "Movie ID does not exist", http.StatusNotFound)
    }
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
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	found := false

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(1000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	if !found {
        http.Error(w, "Movie ID does not exist", http.StatusNotFound)
    }
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Title: "Inception", ISBN: "123456", Director: &Director{Name: "Christopher Nolan"}})
	movies = append(movies, Movie{ID: "2", Title: "The Matrix", ISBN: "654321", Director: &Director{Name: "Lana Wachowski, Lilly Wachowski"}})
	movies = append(movies, Movie{ID: "3", Title: "Interstellar", ISBN: "234567", Director: &Director{Name: "Christopher Nolan"}})
	movies = append(movies, Movie{ID: "4", Title: "The Godfather", ISBN: "765432", Director: &Director{Name: "Francis Ford Coppola"}})
	movies = append(movies, Movie{ID: "5", Title: "The Dark Knight", ISBN: "345678", Director: &Director{Name: "Christopher Nolan"}})
	movies = append(movies, Movie{ID: "6", Title: "Pulp Fiction", ISBN: "876543", Director: &Director{Name: "Quentin Tarantino"}})
	movies = append(movies, Movie{ID: "7", Title: "Fight Club", ISBN: "456789", Director: &Director{Name: "David Fincher"}})
	movies = append(movies, Movie{ID: "8", Title: "Forrest Gump", ISBN: "987654", Director: &Director{Name: "Robert Zemeckis"}})
	movies = append(movies, Movie{ID: "9", Title: "The Shawshank Redemption", ISBN: "567890", Director: &Director{Name: "Frank Darabont"}})
	movies = append(movies, Movie{ID: "10", Title: "The Lord of the Rings: The Return of the King", ISBN: "098765", Director: &Director{Name: "Peter Jackson"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	
	log.Fatal(http.ListenAndServe(":8000", r))
}