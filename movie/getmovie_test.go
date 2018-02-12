package movie_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/Infra-Red/myawesomeapi/movie"
	"github.com/Infra-Red/myawesomeapi/movie/moviefakes"
	"github.com/gorilla/mux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("/movie", func() {
	var (
		fakemovieSource *moviefakes.FakeMovieSource
		testServer      *httptest.Server
		testClient      *http.Client
		baseURL         string
	)

	BeforeEach(func() {
		fakemovieSource = new(moviefakes.FakeMovieSource)
		handler := movie.GetMovieHandler(fakemovieSource)
		server := mux.NewRouter()
		server.Handle("/movie", handler).Methods("GET")
		testServer = httptest.NewServer(server)
		baseURL = testServer.URL
		testClient = http.DefaultClient
	})

	AfterEach(func() {
		testServer.Close()
	})

	Describe("GET", func() {
		var (
			response *http.Response
			err      error
			body     []byte
		)

		BeforeEach(func() {
			fakemovieSource.GetNextMovieStub = func() (string, string, string) {
				return "simple movie", "it's just a movie", "image-data"
			}

			url := fmt.Sprintf("%v/%v", baseURL, "/movie")
			response, err = http.Get(url)

			if err != nil {
				log.Fatal(err)
			}

			body, err = ioutil.ReadAll(response.Body)
			response.Body.Close()

			if err != nil {
				log.Fatal(err)
			}
		})

		It("returns a 200", func() {
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns the movie information from the movieSource", func() {
			Expect(body).To(MatchJSON([]byte(`{"name":"simple movie","description":"it's just a movie","image":"image-data"}`)))
		})
	})
})
