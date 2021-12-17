package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

type Player struct {
    Name string `json:name`
    Nickname string `json:nickname`
}

type Players []Player

func allPlayers(w http.ResponseWriter, r *http.Request) {
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

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Bora pro dojô? obs: isso foi feito em GO")
}

func handleRequests(){
    http.HandleFunc("/", homePage)
    http.HandleFunc("/players", allPlayers)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
    handleRequests()
}
