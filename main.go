package main

import (
	"flag"
	"fmt"
	"os"
	"scraper/scraper"
	"strings"
)

func main() {
  // Flags
  var pdfFlag *bool = flag.Bool("pdf",false,"Convert images to pdf")
  var listComicsFlag *bool = flag.Bool("list-comics",false,"List all available comics")
  // var listIssuesFlag *bool = flag.Bool("list-issues",false,"List all issues of a comic")
  flag.BoolVar(listComicsFlag,"lc",false,"")
  // flag.BoolVar(listIssuesFlag,"li",false,"")
  

  flag.Parse()
	if flag.NArg() != 1 && !*listComicsFlag {
		fmt.Println("Usage:")
		fmt.Println("To download an issue:")
		fmt.Println("program baseURL/comicName/issue")
		fmt.Println()
		fmt.Println("To download all issues of a comic:")
		fmt.Println("program baseURL/comicName")
		fmt.Println()
    fmt.Println("To list all available comics:")
    fmt.Println("program -lc or program --list-comics")
		return
	}

  // Parse argument
	var ComicRoute = strings.TrimPrefix(flag.Arg(0),"https://www.omgbeaupeep.com/comics")
  ComicRoute = strings.TrimSuffix(ComicRoute,"/")

  // Searching
  if *listComicsFlag {
   results,err := scraper.GetComicList()
   if err != nil {
     fmt.Println("Error getting the names of the comics: ",err)
     os.Exit(1)
   }
   for idx,r := range results {
    fmt.Printf("%d - %s\n",idx,r)
   }
   os.Exit(0)
  }

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
