package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/db"
	"github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/poker"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := db.NewInMemoryStore()
	server := poker.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewGetScoreRequest(player))
		poker.AssertStatus(t, response.Code, http.StatusOK)

		poker.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewLeagueRequest())
		poker.AssertStatus(t, response.Code, http.StatusOK)

		got := poker.GetLeagueFromResponse(t, response.Body)
		want := []poker.Player{
			{"Pepper", 3},
		}
		poker.AssertLeague(t, got, want)
	})

}
