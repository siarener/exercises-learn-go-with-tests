package tests

import (
	"log"
	"net/http"

	"github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/db"
	"github.com/apfelkraepfla/exercises-learn-go-with-tests/app/http-server/poker"
)

func main() {
	server := &poker.PlayerServer{Store: db.NewInMemoryStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
