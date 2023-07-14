package main

import (
	"log"
	"net/http"

	"github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/poker"
)

func main() {
	handler := http.HandlerFunc(poker.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
