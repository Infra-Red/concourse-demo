package moviefakes

import (
	"sync"

	"github.com/Infra-Red/myawesomeapi/movie"
)

type FakeMovieSource struct {
	GetNextMovieStub        func() (string, string, string)
	getNextMovieMutex       sync.RWMutex
	getNextMovieArgsForCall []struct{}
	getNextMovieReturns     struct {
		result1 string
		result2 string
		result3 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMovieSource) GetNextMovie() (string, string, string) {
	fake.getNextMovieMutex.Lock()
	fake.getNextMovieArgsForCall = append(fake.getNextMovieArgsForCall, struct{}{})
	fake.recordInvocation("GetNextMovie", []interface{}{})
	fake.getNextMovieMutex.Unlock()
	if fake.GetNextMovieStub != nil {
		return fake.GetNextMovieStub()
	} else {
		return fake.getNextMovieReturns.result1, fake.getNextMovieReturns.result2, fake.getNextMovieReturns.result3
	}
}

func (fake *FakeMovieSource) GetNextMovieCallCount() int {
	fake.getNextMovieMutex.RLock()
	defer fake.getNextMovieMutex.RUnlock()
	return len(fake.getNextMovieArgsForCall)
}

func (fake *FakeMovieSource) GetNextMovieReturns(result1 string, result2 string, result3 string) {
	fake.GetNextMovieStub = nil
	fake.getNextMovieReturns = struct {
		result1 string
		result2 string
		result3 string
	}{result1, result2, result3}
}

func (fake *FakeMovieSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getNextMovieMutex.RLock()
	defer fake.getNextMovieMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMovieSource) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ movie.MovieSource = new(FakeMovieSource)
