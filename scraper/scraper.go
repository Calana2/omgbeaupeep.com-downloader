package scraper

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
  // External 
	"github.com/gocolly/colly"
	"github.com/jung-kurt/gofpdf"
)


// ================================== Global Variables ================================== //

// ------------------------ ANSI Colors ------------------------ //
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"





// ================================== Public Functions ================================== //

// ------------------------ Download Functions ------------------------ //

/*
    Download a single issue.
*/
func DownloadIssue(URL string, toPDF bool) error {
	// Verify if it's a correct URL
	parts := strings.Split(URL, "/comics/")
	if len(parts) < 2 {
		return fmt.Errorf("The URL does not have the correct format.\n")
	}
  
  // Create output directory
	issuePath := parts[1]
	outputPath := filepath.Join("output", issuePath)
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		fmt.Printf(" Creating directory: ./%s/\n", outputPath)
		err = os.MkdirAll(outputPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf(Red+" Error creating the dir: %v\n"+Reset, err)
		}
	}

	// Colly events
  domain,_ := url.Parse(URL)
	c := colly.NewCollector(colly.AllowedDomains(domain.Host))
	c.AllowURLRevisit = true
	c.OnHTML("img[data-omv-prev]", func(e *colly.HTMLElement) {
		imgSrc := e.Attr("src")
    fmt.Println(imgSrc)
    if strings.Contains(URL, "/comics/") {
			filename := filepath.Base(imgSrc)
			imagePath := filepath.Join(outputPath, filename)
			if _, err := os.Stat(imagePath); err == nil {
				fmt.Println(" Image exists:", imgSrc)
				return
			}
      err := downloadImage(domain.Scheme + "://" + domain.Host +imgSrc, imagePath)
			if err != nil {
				fmt.Println(Red+" Error downloading the image: "+Reset, err)
			} else {
				fmt.Println(Green+" Image downloaded:"+Reset, imgSrc)
			}
		}
	})

	// CLI output
	fmt.Println()
	fmt.Println("Downloading Issue:")
	fmt.Println(Green + " " + URL + Reset)
	fmt.Println()

  // Download Images
	err := c.Visit(URL)
	if err != nil {
		fmt.Println(Red+"Error visiting the page: "+Reset, err)
		os.Exit(1)
	}
	index := 2
	for {
		err := c.Visit(URL + "/" + strconv.Itoa(index))
		if err != nil {
			if err.Error() == "Not Found" {
				fmt.Println(Green + " Issue downloaded successfully." + Reset)
				fmt.Println()
				break
			} else {
				fmt.Println(Red+"Error visiting the page: "+Reset, err)
				fmt.Println("Retrying...")
				time.Sleep(1)
				continue
			}
		}
		index++
	}

  // convert to PDF
	if toPDF {
		comicToPDF(outputPath)
	}
	return nil
}





/* 
    Download the entire comic.
*/
func DownloadAllChapters(URL string, toPDF bool) error {
	// Extract subdomain
	domainParts := strings.Split(URL, "//")
	if len(domainParts) < 2 {
		return fmt.Errorf("Invalid URL: couldn't extract domain")
	}
	
	subdomainParts := strings.Split(domainParts[1], ".")
	if len(subdomainParts) < 3 {
		return fmt.Errorf("Invalid URL: couldn't extract subdomain")
	}
	subdomain := subdomainParts[0]

	// Create output directory
	var outputPath string
	if subdomain == "reader" {
		parts := strings.Split(URL, "/comics/")
		if len(parts) < 2 {
			return fmt.Errorf("The URL does not have the correct format")
		}
		outputPath = filepath.Join("output", parts[1])
	} else {
		outputPath = filepath.Join("output", subdomain)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		fmt.Printf("Creating directory: ./%s/\n", outputPath)
		err = os.MkdirAll(outputPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf(Red+"Error creating directory:"+Reset+"%v\n", err)
		}
	}

	// CLI output
	fmt.Println()
	fmt.Println("Downloading Comic:")
	fmt.Println(Green + " " + URL + Reset)
	fmt.Println()

	// Colly collector setup
	var c *colly.Collector
	if subdomain == "reader" {
		c = colly.NewCollector(colly.AllowedDomains("reader.popsensei.com"))
		c.AllowURLRevisit = true
		c.OnHTML("select.change-chapter", func(e *colly.HTMLElement) {
			e.ForEach("option", func(_ int, el *colly.HTMLElement) {
				issue := el.Attr("value")
        fmt.Print(URL)
				if issue != "" {
					err := DownloadIssue(URL+"/"+issue, toPDF)
					if err != nil {
						fmt.Printf(Yellow+"Warning:"+Reset+"Error downloading issue %s: %s\n", issue, err)
					}
				}
			})
		})
	} else {
		domain := fmt.Sprintf("%s.popsensei.com", subdomain)
		c = colly.NewCollector(colly.AllowedDomains(domain))
		c.OnHTML("ul.wp-block-list li a", func(e *colly.HTMLElement) {
			issueURL := e.Attr("href")
			if issueURL != "" {
				err := DownloadIssue(issueURL, toPDF)
				if err != nil {
					fmt.Printf(Yellow+"Warning:"+Reset+"Error downloading issue %s: %s\n", issueURL, err)
				}
			}
		})
	}

	// Action
	err := c.Visit(URL)
	if err != nil {
		fmt.Println(Red+" Error visiting the page: "+Reset, err)
		os.Exit(1)
	}
	return nil
}





// ------------------------ List Functions  ------------------------ //

/*
    Get the list of comics from the site.
*/
func GetComicList() (list []string, err error) {
  // but not for the others
	c := colly.NewCollector(colly.AllowedDomains("reader.popsensei.com"))
	c.OnHTML("ul.wp-block-list li", func(e *colly.HTMLElement) {
		list = append(list, fmt.Sprintf("%-60s %s\n", e.Text, e.ChildAttr("a", "href")))
	})
	err = c.Visit("https://reader.popsensei.com/")
	if err != nil {
		return nil, err
	}
	return list, nil
}





/*
    Get the list of issues of a comic.
    Some comics have their own subdomain and the way they list the chapters it's different.
*/
func GetIssueList(comicURL string) (list []string, err error) {
	subdomain := strings.Split(comicURL, "//")[1]
	subdomain = strings.Split(subdomain, ".popsensei")[0]
	var domain string
	var c *colly.Collector
	if subdomain == "reader" {
		domain = "reader.popsensei.com"
		c = colly.NewCollector(colly.AllowedDomains(domain))
		c.OnHTML("select.change-chapter option", func(e *colly.HTMLElement) {
			cleanText := strings.ReplaceAll(e.Text, "\x09", "")
			cleanText = strings.ReplaceAll(cleanText, "\x0a", "")
      issueNumber := strings.Split(cleanText," ")[0]
			list = append(list, fmt.Sprintf("%-60s %s\n", cleanText, comicURL + "/" + issueNumber ))
		})
	} else {
		// The other way
		domain = fmt.Sprintf("%s.popsensei.com", subdomain)
		c = colly.NewCollector(colly.AllowedDomains(domain))
		c.OnHTML("ul.wp-block-list li", func(e *colly.HTMLElement) {
			list = append(list, fmt.Sprintf("%-60s %s\n", e.Text, e.ChildAttr("a", "href")))
		})
	}
	err = c.Visit(comicURL)
	if err != nil {
		return nil, err
	}
	return list[:len(list)-1], nil
}





// ================================== Private Functions ================================== //

// ------------------------ File Operations ------------------------ //

/*
    Download an image via an HTTP GET request
*/
func downloadImage(url string, outputPath string) error {
	// Download the image
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return fmt.Errorf(response.Status)
	}
	// Save the file
	outfile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outfile.Close()
	_, err = io.Copy(outfile, response.Body)
	return err
}





/*
    Convert the images in a directory to a PDF, then delete the images.
*/
func comicToPDF(imageDirectory string) {
	// Convert images in a directory to PDF
	root := os.DirFS(imageDirectory)
	images, err := fs.Glob(root, "*.jpg")
	if err != nil {
		fmt.Println(Red+"Error:"+Reset+" An error ocurred trying to convert the comic to PDF: ", err)
		return
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pageWidth, pageHeight := pdf.GetPageSize()
	for _, imageName := range images {
		pdf.AddPage()
		pdf.ImageOptions(filepath.Join(imageDirectory, imageName),
			0, 0, pageWidth, pageHeight, false,
			gofpdf.ImageOptions{ImageType: "jpg"}, 0, "")
		fmt.Println(" " + imageName + " added to the PDF.")
	}

	pdfName := strings.Split(imageDirectory, "/")[1] + ".pdf"
	err = pdf.OutputFileAndClose(filepath.Join(imageDirectory, pdfName))
	if err != nil {
		fmt.Println(Red+"Error:"+Reset+" An error ocurred trying to convert the comic to PDF: ", err)
		return
	}
	// Delete images
	for _, imageName := range images {
		os.Remove(filepath.Join(imageDirectory, imageName))
	}
	fmt.Println(Green + " " + pdfName + " created successfully." + Reset)
	fmt.Println()
}





/*
    Determine if a URL from popsensei is a comic URL or an issue URL.
*/
func DetermineComicType(comicURL string) string {
	comicURL = strings.TrimSuffix(comicURL, "/")
	u, err := url.Parse(comicURL)
	if err != nil {
		return "invalid"
	}

	pathParts := strings.Split(u.Path, "/")

	if !strings.HasPrefix(u.Host, "reader.") && len(pathParts) <= 2 {
		return "full_comic"
	}

	if strings.HasPrefix(u.Host, "reader.") {
		if len(pathParts) >= 4 {
			return "single_issue"
		}
		return "full_comic"
	}

	if len(pathParts) >= 4 {
		return "single_issue"
	}
	return "full_comic"
}
