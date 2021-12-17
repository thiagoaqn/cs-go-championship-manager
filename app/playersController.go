package app

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Player struct {
	Name string `json:name`
	Nickname string `json:nickname`
}

type Players []Player

func (app *App) ListPlayersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		players := Players{
			Player{Name:"Thiago", Nickname: "thiagoaqn"},
			Player{Name:"Bruno", Nickname: "brunolrib"},
			Player{Name:"Luiz", Nickname: "LuizNach"},
			Player{Name:"Elkerton", Nickname: "elkerton"},
			Player{Name:"Fernando", Nickname: "ffrm"},
		}

		fmt.Println("Hit in players endpoint")
		json.NewEncoder(w).Encode(players)
	}
}