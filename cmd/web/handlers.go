package main

import (
	"fmt"
	"net/http"
	//"sweater_weather/services"

	"sweater_weather.kyleschulz.net/services"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	w.Write([]byte("Hello from sweater weather"))
}

func (app *application) forecast(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if  location == "" {
		app.notFound(w)
		return
	}

	// Call GetLatLong from the services package
	latLong, err := services.GetLatLong(location)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching location data: %v", err), http.StatusInternalServerError)
		return
	}

	//fmt.Fprintf(w, "Latitude and longitude for %s is %s", location, res)
	fmt.Fprintf(w, "Here's the lat long: %v", latLong)
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	// This will eventuall be a post...
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
