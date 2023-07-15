package db

import (
	"strings"
	"testing"

	"github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/poker"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []poker.Player{
			{Name: "Cleo", Wins: 10},
			{Name: "Chris", Wins: 33},
		}

		poker.AssertLeague(t, got, want)

		got = store.GetLeague()
		poker.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEquals(t, got, want)
	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
