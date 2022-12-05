package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// func main() {
// 	// get the path of the mp4 video
// 	videoPath := os.Args[1]
// 	// get the file name of the video
// 	videoName := filepath.Base(videoPath)
// 	// get the directory of the video
// 	videoDir := filepath.Dir(videoPath)
// 	// create the output file name
// 	outputFileName := videoName[:len(videoName)-4] + ".srt"

// 	// execute the command to generate the subtitles
// 	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vf", "subtitles="+videoDir+"/"+outputFileName)
// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("Subtitles created successfully!")
//}

func main() {
	// get the path of the mp4 video
	videoPath := os.Args[1]
	// get the file name of the video
	videoName := filepath.Base(videoPath)
	// get the directory of the video
	videoDir := filepath.Dir(videoPath)
	// create the output file name
	outputFileName := videoName[:len(videoName)-4] + ".srt"

	// execute the command to generate the subtitles
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vf", "subtitles="+videoDir+"/"+outputFileName)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	// wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Subtitles created successfully!")
}
