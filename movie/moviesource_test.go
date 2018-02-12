package movie_test

import (
	"github.com/Infra-Red/myawesomeapi/movie"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MovieSource", func() {
	var subject movie.MovieSource

	BeforeEach(func() {
		movieFileNames := []string{
			"images/bladerunner.png",
			"images/recall.png",
			"images/interstellar.png",
			"images/chucky.png",
		}
		subject = movie.NewMovieSource(movieFileNames)
	})

	It("returns movie data in a cycle", func() {
		movieName, movieDescription, movieImageData := subject.GetNextMovie()
		Expect(movieName).To(Equal("Blade Runner"))
		Expect(movieDescription).To(Equal("wow, blade runner"))
		Expect(movieImageData).ToNot(BeNil())

		movieName, movieDescription, movieImageData = subject.GetNextMovie()
		Expect(movieName).To(Equal("Total Recall"))
		Expect(movieDescription).To(Equal("hey look, such a great movie"))
		Expect(movieImageData).ToNot(BeNil())

		movieName, movieDescription, movieImageData = subject.GetNextMovie()
		Expect(movieName).To(Equal("Interstellar"))
		Expect(movieDescription).To(Equal("nolan is a genius"))
		Expect(movieImageData).ToNot(BeNil())

		movieName, movieDescription, movieImageData = subject.GetNextMovie()
		Expect(movieName).To(Equal("Chucky"))
		Expect(movieDescription).To(Equal("such a sweet guy"))
		Expect(movieImageData).ToNot(BeNil())

		// Cycles back to the beginning
		movieName, movieDescription, movieImageData = subject.GetNextMovie()
		Expect(movieName).To(Equal("Blade Runner"))
		Expect(movieDescription).To(Equal("wow, blade runner"))
		Expect(movieImageData).ToNot(BeNil())
	})
})
