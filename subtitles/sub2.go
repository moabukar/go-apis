package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Get the video file name from the command line argument
	videoFileName := os.Args[1]

	// Extract the subtitles from the video
	subtitleFileName := strings.Replace(videoFileName, ".mp4", ".srt", 1)
	cmd := exec.Command("ffmpeg", "-i", videoFileName, subtitleFileName)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Convert the subtitles to a txt file
	txtFileName := strings.Replace(videoFileName, ".mp4", ".txt", 1)
	cmd = exec.Command("ffmpeg", "-i", subtitleFileName, txtFileName)
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Subtitles successfully created!")
}
