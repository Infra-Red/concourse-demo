package movie

import (
	"bufio"
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
)

type MovieSource interface {
	GetNextMovie() (string, string, string)
}

func NewMovieSource(movieFileNames []string) MovieSource {
	return &movieSource{
		callCount: 0,
		movieNames: []string{
			"Blade Runner",
			"Total Recall",
			"Interstellar",
			"Chucky",
		},
		movieDescriptions: []string{
			"wow, blade runner",
			"hey look, such a great movie",
			"nolan is a genius",
			"such a sweet guy",
		},
		movieFileNames: movieFileNames,
	}
}

type movieSource struct {
	callCount         int
	movieNames        []string
	movieDescriptions []string
	movieFileNames    []string
}

func (s *movieSource) GetNextMovie() (string, string, string) {
	index := s.callCount % 4
	s.callCount += 1
	movieRelativeFilePath := s.movieFileNames[index]
	movieFilePath, filePathErr := filepath.Abs(movieRelativeFilePath)
	if filePathErr != nil {
		log.Fatal(filePathErr)
	}

	movieFile, err := os.Open(movieFilePath)
	if err != nil {
		log.Fatal(err)
	}

	movieFileInfo, _ := movieFile.Stat()
	var size int64 = movieFileInfo.Size()
	movieFileBase64Bytes := make([]byte, size)
	movieFileReader := bufio.NewReader(movieFile)
	movieFileReader.Read(movieFileBase64Bytes)

	movieFileBase64EncodedString := base64.StdEncoding.EncodeToString(movieFileBase64Bytes)

	movieName := s.movieNames[index]
	movieDescription := s.movieDescriptions[index]
	return movieName, movieDescription, movieFileBase64EncodedString
}
