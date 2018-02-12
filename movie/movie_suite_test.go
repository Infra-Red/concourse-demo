package movie_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMovie(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Movie Suite")
}
