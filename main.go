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
  var helpFlag *bool = flag.Bool("help",false,"")
  var pdfFlag *bool = flag.Bool("pdf",false,"Convert images to pdf")
  var listComicsFlag *bool = flag.Bool("list-comics",false,"List all available comics")
  var listIssuesFlag *bool = flag.Bool("list-issues",false,"List all issues of a comic")
  flag.BoolVar(helpFlag,"h",false,"")
  flag.BoolVar(listComicsFlag,"lc",false,"")
  flag.BoolVar(listIssuesFlag,"li",false,"")
  

  flag.Parse()
	if *helpFlag  || (flag.NArg() != 1 && !*listComicsFlag) {
    fmt.Println("A tool to download comics from https://www.omgbeaupeep.com\n")
		fmt.Println("Usage: omgb [OPTIONS] [comicURL|issueURL]\n")
    fmt.Println("Options:")
    fmt.Println("-h, --help  Print help")
    fmt.Println("-lc, --list-comics  List all available comics")
    fmt.Println("-li, --list-issues <issueURL>  List all issues of a comic")
    fmt.Println("--pdf Convert issues to pdf")
		os.Exit(0)
	}


// Searching
  // comics
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
  // issues
  if *listIssuesFlag {
   results,err := scraper.GetIssueList(flag.Arg(0))
   if err != nil {
     fmt.Println("Error getting the names of the issues: ",err)
     os.Exit(1)
   }
   for _,r := range results {
    fmt.Printf("%s\n",r)
   }
   os.Exit(0)
  }

// Download
	var ComicRoute = strings.TrimPrefix(flag.Arg(0),"https://www.omgbeaupeep.com/comics")
  ComicRoute = strings.TrimSuffix(ComicRoute,"/")
	switch strings.Count(ComicRoute, "/") {
	case 2:
		scraper.DownloadComic(ComicRoute, *pdfFlag)
	case 1:
		scraper.DownloadAllChapters(ComicRoute, *pdfFlag)
	default:
		fmt.Printf("Invalid usage")
	}
}
