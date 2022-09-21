package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gocolly/colly"
)

var scrapedPageNum int

func createDestFileAndWriter(name string) csv.Writer {
	//setting up the file where we store collected data
	fName := filepath.Join("D:\\", "go projects", "cwst go", "CWST-GO", "target folder", name+".csv")
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, error :%q", err)
	}
	defer file.Close()

	//writer that writes the collected data into our file
	writer := csv.NewWriter(file)

	//after the file is written, what it is in the buffer goes in writer and then passed to file
	defer writer.Flush()

	return *writer
}

func parser() *colly.Collector {

	writer := createDestFileAndWriter("index1")

	//collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.publi24.ro"),
	)

	//HTML parser for seller name
	c.OnHTML(".userdata h3", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("a"), //specific tag of the info
		})
	})

	//HTML parser for the date on which the item was posted online
	c.OnHTML(".medium-7", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("i"), //specific tag of the info
		})
	})

	//HTML parser for item title
	c.OnHTML(".large-8", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("h1"), //specific tag of the info
		})
	})

	//HTML parser for item price
	c.OnHTML(".large-4", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("span"), //specific tag of the info
		})
	})

	//HTML parser for item description
	c.OnHTML(".article-detail", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("a"),
			e.ChildText("span"), //specific tag of the info
		})
	})

	//HTML parser for location
	c.OnHTML(".medium-5", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("a"),
		})
	})

	//base64 encoding for all images
	c.OnHTML(".thumbZone img", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			imgtoBase64(e.Attr("src")),
		})
	})

	return c
}

func scrape(webSiteAddr string) {
	cly := parser()
	scrapedPageNum = 1
	fmt.Printf("Scraping page " + strconv.Itoa(scrapedPageNum) + " ... \n")
	cly.Visit(webSiteAddr)
	log.Printf("\n\nScraping Page " + strconv.Itoa(scrapedPageNum) + " Complete\n\n")
	log.Println(cly)
}

// functions for Base 64 image
func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func imgtoBase64(adr string) string {
	resp, err := http.Get(adr)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string
	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(bytes)
	return base64Encoding
}

func main() {
	//setting up the file where we store collected data
	fName := filepath.Join("D:\\", "go projects", "cwst go", "CWST-GO", "target folder", "index1.csv")
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, error :%q", err)
	}
	defer file.Close()

	//writer that writes the collected data into our file
	writer := csv.NewWriter(file)

	//after the file is written, what it is in the buffer goes in writer and then passed to file
	defer writer.Flush()

	//collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.publi24.ro"),
	)

	//HTML parser for seller name
	c.OnHTML(".userdata h3", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("a"), //specific tag of the info
		})
	})

	//HTML parser for the date on which the item was posted online
	c.OnHTML(".medium-7", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("i"), //specific tag of the info
		})
	})

	//HTML parser for item title
	c.OnHTML(".large-8", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("h1"), //specific tag of the info
		})
	})

	//HTML parser for item price
	c.OnHTML(".large-4", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("span"), //specific tag of the info
		})
	})

	//HTML parser for item description
	c.OnHTML(".article-detail", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("a"),
			e.ChildText("span"), //specific tag of the info
		})
	})

	//HTML parser for location
	c.OnHTML(".medium-5", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("a"),
		})
	})

	//base64 encoding for all images
	c.OnHTML(".thumbZone img", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			imgtoBase64(e.Attr("src")),
		})
	})

	//visiting 3 target pages
	fmt.Printf("Scraping page 1 ... \n")
	c.Visit("https://www.publi24.ro/anunturi/auto-moto/masini-second-hand/vw/golf/anunt/vw-golf-6-diesel-euro-5/8248h8271g4875fhd1d60789036fe1h7.html")
	fmt.Printf("Scraping page 2 ... \n")
	c.Visit("https://www.publi24.ro/anunturi/imobiliare/de-vanzare/case/vila/anunt/casa-individuala-cu-garaj-str-dumitru-mocanu/h0g0dg65hffi7467e6h1gg818hiid095.html")
	fmt.Printf("Scraping page 3 ... \n")
	c.Visit("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html")

	log.Printf("\n\nScraping Complete\n\n")
	log.Println(c)

	//scrape("https://www.publi24.ro/anunturi/auto-moto/masini-second-hand/vw/golf/anunt/vw-golf-6-diesel-euro-5/8248h8271g4875fhd1d60789036fe1h7.html")
}
