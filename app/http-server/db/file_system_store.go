package db

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/poker"
)

// FileSystemPlayerStore stores players in the filesystem.
type FileSystemPlayerStore struct {
	Database *json.Encoder
	league   poker.League
}

// NewFileSystemPlayerStore creates a FileSystemPlayerStore initialising the store if needed.
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	file.Seek(0, 0)
	league, err := poker.NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}
	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil

}

// GetLeague returns the scores of all the players.
func (f *FileSystemPlayerStore) GetLeague() poker.League {
	return f.league
}

// GetPlayerScore retrieves a player's score.
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin will store a win for a player, incrementing wins if already known.
func (f *FileSystemPlayerStore) RecordWin(name string) {

	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, poker.Player{Name: name, Wins: 1})
	}

	f.Database.Encode(f.league)
}
