package poker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"

type Player struct {
	Name string
	Wins int
}

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	Store PlayerStore

	/* By embedding `http.Handler``, we give all its methods
	to Player Server, i.e. `ServeHTTP``. Hence, we do not have
	to implement our own `ServeHTTP`` method.

	However, with this we also expose all public methods
	and fields of the embedded type. Because of this, we
	opted to only embed the interface we wanted exposed,
	instead of the concrete type, i.e. `http.ServeMux`.
	If we would have used http.ServeMux, users of
	`PlayerServer` would also be able to add new routes
	to the server by using `Handle`. */
	http.Handler
}

// NewPlayerServer creates a PlayerServer with routing configured.
func NewPlayerServer(store PlayerStore) *PlayerServer {

	p := new(PlayerServer)
	p.Store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	/* Fill in the http.Handler by assigning it to the router.
	This can be done because http.ServeMux has the method ServeHTTP. */
	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.Store.GetLeague())
}

// func (p *PlayerServer) getLeagueTable() []Player {
// 	return []Player{{
// 		Name: "Chris",
// 		Wins: 10},
// 	}
// }

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
