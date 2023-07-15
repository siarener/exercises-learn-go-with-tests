package db

import (
	"io"

	"github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/poker"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

// Read JSON from Reader
func (f *FileSystemPlayerStore) GetLeague() []poker.Player {
	f.database.Seek(0, 0)
	league, _ := poker.NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	var wins int

	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}
