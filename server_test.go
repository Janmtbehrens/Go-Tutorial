package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type TestPlayerData struct {
	Name  string
	Score string
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) IncPlayerScore(name string) {
	s.scores[name] = s.scores[name] + 1
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  -15,
			"Jerome": 2,
			"Pilly":  -0,
		},
	}

	playerServer := NewPlayerServer(&store)

	t.Run("Returns Score", func(t *testing.T) {
		for name, score := range store.scores {
			request := newGetScoreRequest(name)
			response := httptest.NewRecorder()

			playerServer.ServeHTTP(response, request)

			got := response.Body.String()
			want := strconv.Itoa(score)

			assertResponseBody(t, got, want)
		}
	})

	t.Run("404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("nil")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		assertResponseBodyStatus(t, got, want)
	})

	t.Run("200 on found player", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusOK

		assertResponseBodyStatus(t, got, want)
	})
}

func TestStoreScore(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
		},
	}

	playerServer := NewPlayerServer(&store)

	t.Run("Adding Score is recorded", func(t *testing.T) {
		request := newPostScoreRequest("Pepper")
		response := httptest.NewRecorder()
		// Inc Score
		playerServer.ServeHTTP(response, request)

		// Check Score
		request = newGetScoreRequest("Pepper")
		playerServer.ServeHTTP(response, request)

		got := response.Body.String()
		want := "21"

		assertResponseBody(t, got, want)
	})
	t.Run("Accepted status on POST", func(t *testing.T) {
		request := newPostScoreRequest("Pepper")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusAccepted

		assertResponseBodyStatus(t, got, want)
	})
}

func TestLeague(t *testing.T){
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)

	t.Run("Returns 200 on /league", func(t *testing.T){
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBodyStatus(t, response.Code, http.StatusOK)
	})
}

func newPostScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/players/"+name, nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertResponseBodyStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %d want %d", got, want)
	}
}
