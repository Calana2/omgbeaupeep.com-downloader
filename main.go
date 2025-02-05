package main

import (
	"fmt"
	"os"
	"scraper/scraper"
	"strings"
//"github.com/jung-kurt/gofpdf"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("To download an issue:")
		fmt.Println("program baseURL/comicName/issue")
		fmt.Println("To download all issues of a comic:")
		fmt.Println("program baseURL/comicName")

		return
	}
  // Parse argument
	var ComicRoute = strings.TrimPrefix(os.Args[1],"https://www.omgbeaupeep.com/comics")
  ComicRoute = strings.TrimSuffix(ComicRoute,"/")
  // Operations
	switch strings.Count(ComicRoute, "/") {
	case 2:
		scraper.DownloadComic(ComicRoute)
	case 1:
		scraper.DownloadAllChapters(ComicRoute)
	default:
		fmt.Printf("Invalid usage")
	}
}
