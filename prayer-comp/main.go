package main

import (
	"encoding/json"
	"net/http"
)

// PrayerTimes model
type PrayerTimes struct {
	Fajr    string `json:"fajr"`
	Dhuhr   string `json:"dhuhr"`
	Asr     string `json:"asr"`
	Maghrib string `json:"maghrib"`
	Isha    string `json:"isha"`
}

func main() {
	http.HandleFunc("/api/prayer-times", prayerTimesHandler)
	http.ListenAndServe(":8000", nil)
}

func prayerTimesHandler(w http.ResponseWriter, r *http.Request) {
	// set prayer times
	prayerTimes := PrayerTimes{
		Fajr:    "5:00am",
		Dhuhr:   "12:00pm",
		Asr:     "3:00pm",
		Maghrib: "5:00pm",
		Isha:    "7:00pm",
	}

	// write response
	response, err := json.Marshal(prayerTimes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
