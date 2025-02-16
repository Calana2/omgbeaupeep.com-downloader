package scraper

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/jung-kurt/gofpdf"
)

// +------------------+
// | Global Variables |
// +------------------+
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

// +------------------+
// | Public Functions |
// +------------------+
func DownloadComic(route string, toPDF bool) {
	// Create output directory
	outputPath := filepath.Join("output", route)
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		fmt.Printf(" Creating directory: ./%s/\n", outputPath)
		err = os.MkdirAll(outputPath, os.ModePerm)
		if err != nil {
			fmt.Printf(Red+" Error creating the dir: %v\n"+Reset, err)
			return
		}
	}
	// Colly events
	c := colly.NewCollector(colly.AllowedDomains("www.omgbeaupeep.com"))
	c.AllowURLRevisit = true
	c.OnHTML("img", func(e *colly.HTMLElement) {
		imgSrc := e.Attr("src")
		if !strings.Contains(imgSrc, "https://www.omgbeaupeep.com/") {
			filename := filepath.Base(imgSrc)
			imagePath := filepath.Join(outputPath, filename)
			if _, err := os.Stat(imagePath); err == nil {
				fmt.Println(" Image exists:", imgSrc)
				return
			}
			err := downloadImage(imgSrc, imagePath)
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
	fmt.Println(Green + " " + route + Reset)
	fmt.Println()
	// Actions
	err := c.Visit("https://www.omgbeaupeep.com/comics" + route)
	if err != nil {
		fmt.Println(Red+"Error visiting the page: "+Reset, err)
		os.Exit(1)
	}
	index := 2
	for {
		err := c.Visit("https://www.omgbeaupeep.com/comics" + route + "/" + strconv.Itoa(index))
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

	if toPDF {
		comicToPDF(outputPath)
	}
}

func DownloadAllChapters(comic string, toPDF bool) {
	// Create output directory
	outputPath := filepath.Join("output", comic)
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		fmt.Printf("Creating directory: ./%s/\n", outputPath)
		err = os.MkdirAll(outputPath, os.ModePerm)
		if err != nil {
			fmt.Printf(Red+"Error creating the dir:"+Reset+"%v\n", err)
			return
		}
	}
	// Colly events
	var firstSelect bool = true
	c := colly.NewCollector(colly.AllowedDomains("www.omgbeaupeep.com"))
	c.AllowURLRevisit = true
	c.OnHTML("select.change-chapter", func(e *colly.HTMLElement) {
		if !firstSelect {
			return
		}
		e.ForEach("option", func(_ int, el *colly.HTMLElement) {
			issue := el.Attr("value")
			if issue != "" {
				DownloadComic(comic+"/"+issue, toPDF)
			} else {
				fmt.Print(Yellow + " Warning:" + Reset + " <option> doesn't have \"value\" attribute")
			}
			firstSelect = false
		})
	})
	// CLI output
	fmt.Println()
	fmt.Println("Downloading Comic:")
	fmt.Println(Green + " " + comic + Reset)
	fmt.Println()
	// Actions
	err := c.Visit("https://www.omgbeaupeep.com/comics" + comic)
	if err != nil {
		fmt.Println(Red+" Error visiting the page: "+Reset, err)
		os.Exit(1)
	}
}

func GetComicList() (list []string, err error) {
	c := colly.NewCollector(colly.AllowedDomains("www.omgbeaupeep.com"))
	c.OnHTML("select.change-manga option[value]", func(e *colly.HTMLElement) {
		list = append(list, e.Text+" (https://www.omgbeaupeep.com/comics/"+e.Attr("value")+")")
	})
	err = c.Visit("https://www.omgbeaupeep.com/comics")
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetIssueList(comicURL string) (list []string, err error) {
	var firstSelect bool = true
	c := colly.NewCollector(colly.AllowedDomains("www.omgbeaupeep.com"))

	c.OnHTML("select.change-chapter", func(e *colly.HTMLElement) {
		if !firstSelect {
			return
		}
		e.ForEach("option", func(_ int, el *colly.HTMLElement) {
			value := el.Attr("value")
			name := strings.ReplaceAll(el.Text[1:], "\t", "")
			list = append(list, name+" ("+comicURL+value+")")
		})
		firstSelect = false
	})
	err = c.Visit(comicURL)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// +-------------------+
// | Private functions |
// +-------------------+
func downloadImage(url string, outputPath string) error {
	// Download the image
	response, err := http.Get("https://www.omgbeaupeep.com" + url)
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
