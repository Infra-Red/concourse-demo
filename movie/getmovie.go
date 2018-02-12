package movie

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetMovieHandler(movieSource MovieSource) http.Handler {
	return &movieHandler{
		movieSource: movieSource,
	}
}

type movieHandler struct {
	movieIndex  int
	movieSource MovieSource
}

type movie struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Base64Image string `json:"image"`
}

func (h *movieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	movieName, movieDescription, movieImageBase64DataString := h.movieSource.GetNextMovie()
	movie := &movie{
		Name:        movieName,
		Description: movieDescription,
		Base64Image: movieImageBase64DataString,
	}

	movieBytes, jsonErr := json.Marshal(movie)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(movieBytes)
}
