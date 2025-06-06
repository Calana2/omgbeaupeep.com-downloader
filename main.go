package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"scraper/scraper"
	"strings"
	"syscall"
)

func init() {
	// Ignore SIGPIPE
	signal.Notify(make(chan os.Signal, 1), syscall.SIGPIPE)
}

func main() {
	// Flags
	var helpFlag *bool = flag.Bool("help", false, "")
	var pdfFlag *bool = flag.Bool("pdf", false, "Convert images to pdf")
	var listComicsFlag *bool = flag.Bool("list-comics", false, "List all available comics")
	var listIssuesFlag *bool = flag.Bool("list-issues", false, "List all issues of a comic")
	flag.BoolVar(helpFlag, "h", false, "")
	flag.BoolVar(listComicsFlag, "lc", false, "")
	flag.BoolVar(listIssuesFlag, "li", false, "")

	flag.Parse()
	if *helpFlag || (flag.NArg() != 1 && !*listComicsFlag) {
		fmt.Println("A tool to download comics from https://www.omgbeaupeep.com\n")
		fmt.Println("Usage: omgb [OPTIONS] [comicURL|issueURL]\n")
		fmt.Println("Options:")
		fmt.Println("-h, --help  Print help")
		fmt.Println("-lc, --list-comics  List all available comics")
		fmt.Println("-li, --list-issues [issueURL]  List all issues of a comic")
		fmt.Println("--pdf Convert issues to pdf")
		os.Exit(0)
	}

	// Searching
	// comics
	if *listComicsFlag {
		results, err := scraper.GetComicList()
		if err != nil {
			fmt.Println("Error getting the names of the comics: ", err)
			os.Exit(1)
		}
		fmt.Println()
		for idx, r := range results {
			fmt.Printf("%d - %s\n", idx, r)
		}
		os.Exit(0)
	}
	// issues
	if *listIssuesFlag {
    normalizedURL := strings.TrimRight(flag.Arg(0),"/")
		results, err := scraper.GetIssueList(normalizedURL)
		if err != nil {
			fmt.Println("Error getting the names of the issues: ", err)
			os.Exit(1)
		}
		fmt.Println()
		for _, r := range results {
			fmt.Printf("%s\n", r)
		}
		os.Exit(0)
	}

	// Download
	var ComicRoute = strings.TrimSpace(flag.Arg(0))
  ComicRoute = strings.TrimRight(ComicRoute,"/")
  var err error;

	switch scraper.DetermineComicType(ComicRoute) {
	case "full_comic":
		scraper.DownloadAllChapters(ComicRoute, *pdfFlag)
	case "single_issue":
    err = scraper.DownloadIssue(ComicRoute, *pdfFlag)
	default:
    fmt.Println("Invalid URL")
		os.Exit(1)
	}
  if err != nil {
      fmt.Println(err)
  }
}
