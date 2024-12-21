package scraper

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"github.com/gocolly/colly"
)

func downloadImage(url string, outputDir string) error {
 // Download the image
 response, err := http.Get("https://www.omgbeaupeep.com"+url)
 if err != nil {
  return err
 }
 defer response.Body.Close()
 if response.StatusCode != 200 {
   return fmt.Errorf(response.Status)
 }
 // Save the file
 filename := filepath.Base(url)
 outputPath := filepath.Join(outputDir,filename)
 outfile, err := os.Create(outputPath)
 if err != nil {
  return err
 }
 defer outfile.Close()
 _, err = io.Copy(outfile,response.Body)
 return err
}



func DownloadComic(route string) {
 // Create output directory
 outputPath := filepath.Join("output",route)
 if _, err := os.Stat(outputPath); os.IsNotExist(err) {
  fmt.Printf("Creating %s\n",outputPath)
  err = os.MkdirAll(outputPath,os.ModePerm)
  if err != nil {
   fmt.Printf("Error creating the dir: %v\n", err)
   return
  }
 }
 // Colly events
  c := colly.NewCollector(colly.AllowedDomains("www.omgbeaupeep.com"),)
  c.OnHTML("img", func(e *colly.HTMLElement) {
    imgSrc := e.Attr("src")
    if !strings.Contains(imgSrc,"https://www.omgbeaupeep.com/") {
     err := downloadImage(imgSrc,outputPath)
     if err != nil {
      fmt.Println("Error downloading the image:", err)
     } else {
      fmt.Println("Image downloaded:", imgSrc)
     }
    }
  })
 // Actions
  fmt.Println("Searching for " + route)
  err := c.Visit("https://www.omgbeaupeep.com/comics/" + route)
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
     continue
    }
   }
   index++ 
  }
}

