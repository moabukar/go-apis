package main

// Imports
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Struct to store prayer times
type PrayerTime struct {
	Fajr    string `json:"Fajr"`
	Dhuhr   string `json:"Dhuhr"`
	Asr     string `json:"Asr"`
	Maghrib string `json:"Maghrib"`
	Isha    string `json:"Isha"`
}

// Handler function to get prayer times
func getPrayerTimes(w http.ResponseWriter, r *http.Request) {
	// Get prayer times from API
	response, err := http.Get("http://api.aladhan.com/v1/timingsByCity?city=London&country=United Kingdom&method=2")

	if err != nil {
		fmt.Printf("The API request failed with error %s\n", err)
		os.Exit(1)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		// Unmarshal the JSON data into our struct
		var prayerTime PrayerTime
		json.Unmarshal(data, &prayerTime)

		// Return the prayer times
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(prayerTime)
	}
}

// Main function
func main() {
	// Create the handler
	http.HandleFunc("/getPrayerTimes", getPrayerTimes)

	// Start the server
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
