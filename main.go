package main

import (
	"fmt"
	"os"
	"scraper/scraper"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("To download an issue:")
		fmt.Println("program /comicName/Issue")
		fmt.Println("To download all issues of a comic:")
		fmt.Println("program /comicName")
		return
	}

	var ComicRoute = os.Args[1]
	switch strings.Count(ComicRoute, "/") {
	case 2:
		scraper.DownloadComic(ComicRoute)
	case 1:
		scraper.DownloadAllChapters(ComicRoute)
	default:
		fmt.Printf("Invalid usage")
	}
}
