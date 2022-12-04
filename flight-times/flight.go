package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FlightTimeResponse is response model
type FlightTimeResponse struct {
	FlightNumber string `json:"flightNumber"`
	Departure    string `json:"departure"`
	Arrival      string `json:"arrival"`
}

// FlightTimes fetches the flight times from Heathrow
func FlightTimes() ([]FlightTimeResponse, error) {
	url := "https://api.heathrow.com/flights/departures/LHR"

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []FlightTimeResponse{}, err
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	client := &http.Client{}

	// Send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		return []FlightTimeResponse{}, err
	}

	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var responseData []FlightTimeResponse

	// Read the data into a byte array
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []FlightTimeResponse{}, err
	}

	// Unmarshal the JSON array data
	if err := json.Unmarshal(body, &responseData); err != nil {
		return []FlightTimeResponse{}, err
	}

	return responseData, nil
}

func main() {
	response, err := FlightTimes()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, flightTime := range response {
		fmt.Printf("%s\t%s\t%s\n", flightTime.FlightNumber, flightTime.Departure, flightTime.Arrival)
	}
}
