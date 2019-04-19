package main

import (
	"repo/mapnificentGenerator"
	"io"
	"os"
	"fmt"
	"net/http"
)

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	
	return err
}

func main() {
	gtfsUrl := "https://api.transitfeeds.com/v1/getLatestFeedVersion?key=d137fd64-dd1f-4b10-aec3-646458732214&feed=sfmta%2F942" // use your url here
	//mapnificentGenerator.Generate("gtfs.zip", "gtfs.bin")
	err := downloadFile("gtfs.zip", gtfsUrl)
	if err != nil {
		fmt.Println(err)
	} else {
		// Send to mapnificentGenerator
	mapnificentGenerator.Generate("gtfs.zip", "bayarea.bin")
	}
}