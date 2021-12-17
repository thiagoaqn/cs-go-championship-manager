package app

import (
	"fmt"
	"net/http"
)

func (app *App) HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Bora pro dojô? obs: isso foi feito em GO")
	}
}