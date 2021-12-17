package main

import (
	"cs-go-championship-manager/app"
	"log"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		log.Println("Server crashed")
	}
}

func main() {
	app := app.New()

	log.Println("Server is Running!")

	http.HandleFunc("/", app.Router.ServeHTTP)

	err := http.ListenAndServe(":8081", nil)

	checkError(err)
}
