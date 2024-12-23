package main

import (
 "os"
 "scraper/scraper"
)

func main() { 
  ComicRoute := "/Avatar_The_Last_Airbender/001/"
  if len(os.Args) == 2 {
    ComicRoute = os.Args[1]
  }

  scraper.DownloadComic(ComicRoute)
}
