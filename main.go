package main

import "scraper/scraper"

const ComicRoute = "/The_Sandman/001/"

func main() { 
  scraper.DownloadComic(ComicRoute)
}
