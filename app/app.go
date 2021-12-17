package app

import (
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func New() *App {
	app := &App{
		Router: mux.NewRouter(),
	}

	app.initRoutes()

	return app
}

func (app *App) initRoutes() {
	app.Router.HandleFunc("/", app.HomeHandler()).Methods("GET")
	app.Router.HandleFunc("/players", app.ListPlayersHandler()).Methods("GET")
}