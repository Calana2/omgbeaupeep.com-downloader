package main

import (
	"fmt"
	"os"
	"scraper/scraper"
)

func main() { 
  if len(os.Args) != 2 {
    fmt.Printf("Usage: program /comicName/Issue/\n")
    return
  }

  var ComicRoute = os.Args[1]
  scraper.DownloadComic(ComicRoute)
}
