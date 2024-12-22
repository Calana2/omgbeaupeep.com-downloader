package scraper

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"github.com/gocolly/colly"
)

func downloadImage(url string, outputPath string) error {
	// Download the image
	response, err := http.Get("https://www.omgbeaupeep.com" + url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return fmt.Errorf(response.Status)
	}
	// Save the file
	outfile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outfile.Close()
	_, err = io.Copy(outfile, response.Body)
	return err
}

func DownloadComic(route string) {
	// Create output directory
	outputPath := filepath.Join("output", route)
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		fmt.Printf("Creating directory: ./%s/\n", outputPath)
		err = os.MkdirAll(outputPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating the dir: %v\n", err)
			return
		}
	}
	// Colly events
	c := colly.NewCollector(colly.AllowedDomains("www.omgbeaupeep.com"))
	c.AllowURLRevisit = true
	c.OnHTML("img", func(e *colly.HTMLElement) {
		imgSrc := e.Attr("src")
		if !strings.Contains(imgSrc, "https://www.omgbeaupeep.com/") {
			filename := filepath.Base(imgSrc)
			imagePath := filepath.Join(outputPath, filename)
      if _,err := os.Stat(imagePath); err == nil {
			fmt.Println("Image exists:", imgSrc)
       return
      }
			err := downloadImage(imgSrc, imagePath)
			if err != nil {
				fmt.Println("Error downloading the image:", err)
			} else {
				fmt.Println("Image downloaded:", imgSrc)
			}
		}
	})
	// Actions
	fmt.Println("Starting task: Download " + route)
	err := c.Visit("https://www.omgbeaupeep.com/comics" + route)
	if err != nil {
		fmt.Println("Error visiting the page: ", err)
		os.Exit(1)
	}
	index := 2
	for {
		err := c.Visit("https://www.omgbeaupeep.com/comics" + route + "/" + strconv.Itoa(index))
		if err != nil {
			if err.Error() == "Not Found" {
				fmt.Println("Task completed.")
				break
			} else {
				fmt.Println("Error visiting the page: ", err)
		                fmt.Println("Retrying...")
				time.Sleep(1)
				continue
			}
		}
		index++
	}
}
