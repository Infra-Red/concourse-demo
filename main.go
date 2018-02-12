package main

import (
	"log"
	"net/http"

	"github.com/Infra-Red/myawesomeapi/movie"
)

func main() {
	movieFileNames := []string{
		"movie/images/bladerunner.png",
		"movie/images/recall.png",
		"movie/images/interstellar.png",
		"movie/images/chucky.png",
	}
	http.Handle("/movie", movie.GetMovieHandler(movie.NewMovieSource(movieFileNames)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
