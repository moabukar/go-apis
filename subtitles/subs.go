package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Function that takes in a video file and creates subtitles
func createSubtitles(video string) {
	// Read the video file
	videoData, err := ioutil.ReadFile(video)
	if err != nil {
		fmt.Println("Error opening video file:", err)
		return
	}

	// Split the video data into lines
	lines := strings.Split(string(videoData), "\n")

	// Create an empty string for the subtitles
	subtitles := ""

	// Loop through the lines of the video and create subtitles
	for _, line := range lines {
		// Check if the line contains the correct format for start and end times
		if !strings.Contains(line, "-") {
			fmt.Println("Error: incorrect format for start and end times")
			return
		}

		// Parse the line to get the start and end times
		startTime, endTime := parseLine(line)

		// Create a new line in the subtitles
		subtitles += startTime + " --> " + endTime + "\n"

		// Get the text for the subtitle
		text := getText(line)

		// Check if the line contains the text for the subtitle
		if text == "" {
			fmt.Println("Error: no text for subtitle")
			return
		}

		// Add the text to the subtitle
		subtitles += text + "\n\n"
	}

	// Write the subtitles to file
	err = ioutil.WriteFile("subtitles.srt", []byte(subtitles), 0644)
	if err != nil {
		fmt.Println("Error writing subtitles to file:", err)
		return
	}
}

// Function that parses a line and returns the start and end times
func parseLine(line string) (string, string) {
	// Split the line to get the start and end times
	times := strings.Split(line, "-")

	// Return the start and end times as strings
	return times[0], times[1]
}

// Function that gets the text for a subtitle
func getText(line string) string {
	// Split the line to get the subtitle text
	text := strings.Split(line, ":")[1]

	// Return the text
	return strings.TrimSpace(text)
}

func main() {
	// Create subtitles for the video
	createSubtitles("mvs.mp4")
}
