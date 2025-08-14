package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	testPlayerName := "Pepper"
	t.Run("Returns Score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "players/"+testPlayerName, nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
