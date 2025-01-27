package services


import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"sweater_weather.kyleschulz.net/internal/models"
)

// create a function that will accept an address
// make a call to https://www.mapquestapi.com/geocoding/v1/address?location=cincinatti,oh
// return the latitude and longitude

//func getLatLong(city string) LatLong {
// https://maps.googleapis.com/maps/api/geocode/json?address=1600+Amphitheatre+Parkway,+Mountain+View,+CA&key=YOUR_API_KEY
func GetLatLong(city string) (models.LatLong, error) {
	latLong := models.LatLong{}

	secrets, err := GetSecrets()

	if err != nil {
		return latLong, fmt.Errorf("Error with GetSecrets(): %v", err)
	}
	fmt.Printf("Secrets: %+v\n", secrets)
	// Use the APIKey from the secrets
	fmt.Println("API Key:", secrets.APIKey)

	// This will call the endpoint and return a lat/long response
	baseURL := "https://maps.googleapis.com/maps/api/geocode/json"
	location := strings.Replace(city, ",", ",+", 1)

	// Parse the base URL
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return latLong, fmt.Errorf("invalid base URL: %w", err) // Return an empty string and the error
	}

	// Add the location as a query parameter
	query := parsedURL.Query()
	query.Set("address", location)
	query.Set("key", secrets.APIKey)
	parsedURL.RawQuery = query.Encode()

	// Create the GET request
	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return latLong, fmt.Errorf("error creating request: %w", err) // Return an empty string and the error
	}

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return latLong, fmt.Errorf("error making HTTP request: %w", err) // Return an empty string and the error
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return latLong, fmt.Errorf("received non-2xx response: %d %s", resp.StatusCode, resp.Status) // Return an empty string and the error
	}

// Read the response body into a byte slice
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return latLong, fmt.Errorf("error reading response body: %w", err)
	}

// Parse the raw response into a map to extract only the lat and lng
	var responseMap map[string]interface{}
	err = json.Unmarshal(responseData, &responseMap)
	if err != nil {
		return models.LatLong{}, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// Navigate through the map to find the lat and lng
	results, ok := responseMap["results"].([]interface{})
	if !ok || len(results) == 0 {
		return latLong, fmt.Errorf("no results found for city: %s", city)
	}

	// Extract latitude and longitude from the first result
	geometry, ok := results[0].(map[string]interface{})["geometry"].(map[string]interface{})
	if !ok {
		return latLong, fmt.Errorf("geometry data not found in response")
	}

	locationData, ok := geometry["location"].(map[string]interface{})
	if !ok {
		return latLong, fmt.Errorf("location data not found in response")
	}

	lat, ok := locationData["lat"].(float64)
	if !ok {
		return latLong, fmt.Errorf("latitude not found in response")
	}

	lng, ok := locationData["lng"].(float64)
	if !ok {
		return latLong, fmt.Errorf("longitude not found in response")
	}

	// Convert lat and lng to string and return them in the LatLong struct
	return models.LatLong{
		Latitude:  fmt.Sprintf("%f", lat),
		Longitude: fmt.Sprintf("%f", lng),
	}, fmt.Errorf("no results found for city: %s", city)
}
