package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/api/v1/forecast", app.forecast)
	mux.HandleFunc("/api/users", app.signup)

	return mux
}
