package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// youtubeDl downloads the specified YouTube video.
func youtubeDl(videoID string, quality string) error {
	// construct the YouTube URL for the video
	youtubeURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)

	// download the video using youtube-dl
	cmd := exec.Command("youtube-dl", "-f", quality, "-o", "%(title)s.%(ext)s", youtubeURL)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// youtubeThumb downloads the thumbnail image for the specified YouTube video.
func youtubeThumb(videoID string) error {
	// construct the URL for the thumbnail image
	thumbURL := fmt.Sprintf("https://img.youtube.com/vi/%s/0.jpg", videoID)

	// download the thumbnail image
	resp, err := http.Get(thumbURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create a file to store the image
	file, err := os.Create("thumb.jpg")
	if err != nil {
		return err
	}
	defer file.Close()

	// copy the image data to the file
	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	// parse the command-line flags
	videoID := flag.String("id", "", "YouTube video ID")
	quality := flag.String("quality", "best", "video quality")
	flag.Parse()

	// validate the flags
	if *videoID == "" {
		log.Fatal("missing YouTube video ID")
	}
	if !strings.HasPrefix(*quality, "best") && !strings.HasPrefix(*quality, "worst") {
		log.Fatal("invalid quality: must be 'best' or 'worst'")
	}

	// download the YouTube video and thumbnail image
	if err := youtubeDl(*videoID, *quality); err != nil {
		log.Fatal(err)
	}
	if err := youtubeThumb(*videoID); err != nil {
		log.Fatal(err)
	}
}
