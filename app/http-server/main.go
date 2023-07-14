package main

import (
	"log"
	"net/http"

	"github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/poker"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &poker.PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
