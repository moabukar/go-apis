package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Get the filename from the command line
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}
	filename := os.Args[1]

	// Read the file
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Split the file into lines
	lines := strings.Split(string(bytes), "\n")

	// Create the subtitles
	subtitles := ""
	for _, line := range lines {
		subtitles += fmt.Sprintf("00:00:00,000 --> 00:00:00,000\n%s\n\n", line)
	}

	// Write the subtitles to a file
	err = ioutil.WriteFile("subtitles.srt", []byte(subtitles), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}

	fmt.Println("Subtitles created successfully!")
}
