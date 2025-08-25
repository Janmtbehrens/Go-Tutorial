package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

type TestPlayerData struct {
	Name  string
	Score string
}

type StubPlayerStore struct {
	scores map[string]int
	league []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) IncPlayerScore(name string) {
	s.scores[name] = s.scores[name] + 1
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  -15,
			"Jerome": 2,
			"Pilly":  -0,
		}, nil,
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
		}, nil,
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

func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)

	t.Run("Returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}

		assertResponseBodyStatus(t, response.Code, http.StatusOK)
	})
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assertResponseBodyStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)

		assertContentType(t, response, jsonContentType)
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
func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}
func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}
