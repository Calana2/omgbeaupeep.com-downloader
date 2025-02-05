package main

import (
	"flag"
	"fmt"
	"scraper/scraper"
	"strings"
)

func main() {
  // Flags
  var pdfFlag *bool = flag.Bool("pdf",false,"Convert images to pdf")
  
  flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage:")
		fmt.Println("To download an issue:")
		fmt.Println("program baseURL/comicName/issue")
		fmt.Println("To download all issues of a comic:")
		fmt.Println("program baseURL/comicName")
		return
	}

  // Parse argument
	var ComicRoute = strings.TrimPrefix(flag.Arg(0),"https://www.omgbeaupeep.com/comics")
  ComicRoute = strings.TrimSuffix(ComicRoute,"/")
  // Operations
	switch strings.Count(ComicRoute, "/") {
	case 2:
		scraper.DownloadComic(ComicRoute, *pdfFlag)
	case 1:
		scraper.DownloadAllChapters(ComicRoute, *pdfFlag)
	default:
		fmt.Printf("Invalid usage")
	}
}
