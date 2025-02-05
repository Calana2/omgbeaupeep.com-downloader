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


// ANSI COLORS
var Reset = "\033[0m" 
var Red = "\033[31m" 
var Green = "\033[32m" 
var Yellow = "\033[33m" 
var Blue = "\033[34m" 
var Magenta = "\033[35m" 
var Cyan = "\033[36m" 
var Gray = "\033[37m" 
var White = "\033[97m"

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
		fmt.Printf(" Creating directory: ./%s/\n", outputPath)
		err = os.MkdirAll(outputPath, os.ModePerm)
		if err != nil {
			fmt.Printf(Red + " Error creating the dir: %v\n" + Reset, err)
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
			if _, err := os.Stat(imagePath); err == nil {
				fmt.Println(" Image exists:", imgSrc)
				return
			}
			err := downloadImage(imgSrc, imagePath)
			if err != nil {
				fmt.Println(Red + " Error downloading the image: " + Reset, err)
			} else {
				fmt.Println(Green + " Image downloaded:" + Reset, imgSrc)
			}
		}
	})
	// CLI output
  fmt.Println()
  fmt.Println("Downloading Issue:")
	fmt.Println(Green + " " + route + Reset)
  fmt.Println()
	// Actions
	err := c.Visit("https://www.omgbeaupeep.com/comics" + route)
	if err != nil {
		fmt.Println(Red + "Error visiting the page: " + Reset, err)
		os.Exit(1)
	}
	index := 2
	for {
		err := c.Visit("https://www.omgbeaupeep.com/comics" + route + "/" + strconv.Itoa(index))
		if err != nil {
			if err.Error() == "Not Found" {
				fmt.Println(Green + "Issue downloaded successfully.")
				fmt.Println()
				break
			} else {
				fmt.Println(Red + "Error visiting the page: " + Reset, err)
				fmt.Println("Retrying...")
				time.Sleep(1)
				continue
			}
		}
		index++
	}
}

func DownloadAllChapters(comic string) {
	// Create output directory
	outputPath := filepath.Join("output", comic)
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		fmt.Printf("Creating directory: ./%s/\n", outputPath)
		err = os.MkdirAll(outputPath, os.ModePerm)
		if err != nil {
			fmt.Printf(Red + "Error creating the dir:" + Reset + "%v\n", err)
			return
		}
	}
	// Colly events
	c := colly.NewCollector(colly.AllowedDomains("www.omgbeaupeep.com"))
	c.AllowURLRevisit = true
	c.OnHTML("select.change-chapter option[value]", func(e *colly.HTMLElement) {
		issue := e.Attr("value")
		if issue != "" {
			DownloadComic(comic + "/" + issue)
		} else {
      fmt.Print(Yellow + " Warning:" + Reset + " <option> doesn't have \"value\" attribute")
		}
	})
	// CLI output
  fmt.Println()
  fmt.Println("Downloading Comic:")
	fmt.Println(Green + " " + comic + Reset)
  fmt.Println()
  // Actions
	err := c.Visit("https://www.omgbeaupeep.com/comics" + comic)
	if err != nil {
		fmt.Println(Red + " Error visiting the page: " + Reset, err)
		os.Exit(1)
	}
}

func ComicToPDF(comic string) {

}
