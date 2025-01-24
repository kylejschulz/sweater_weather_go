package services


import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// create a function that will accept an address
// make a call to https://www.mapquestapi.com/geocoding/v1/address?location=cincinatti,oh
// return the latitude and longitude

//func getLatLong(city string) LatLong {
// https://maps.googleapis.com/maps/api/geocode/json?address=1600+Amphitheatre+Parkway,+Mountain+View,+CA&key=YOUR_API_KEY
func GetLatLong(city string) (string, error) {
	// This will call the endpoint and return a lat/long response
	apiKey := "AIzaSyCb0s-R-lKfq6SAn3eegeSgYXuNN2PzD-k"
	baseURL := "https://maps.googleapis.com/maps/api/geocode/json"
	location := strings.Replace(city, ",", ",+", 1)

	// Parse the base URL
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err) // Return an empty string and the error
	}

	// Add the location as a query parameter
	query := parsedURL.Query()
	query.Set("address", location)
	query.Set("key", apiKey)
	parsedURL.RawQuery = query.Encode()

	// Create the GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err) // Return an empty string and the error
	}

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making HTTP request: %w", err) // Return an empty string and the error
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("received non-2xx response: %d %s", resp.StatusCode, resp.Status) // Return an empty string and the error
	}

	// Read and return the response body
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err) // Return an empty string and the error
	}
	return string(responseData), nil // Return the response data as a string and no error
}
