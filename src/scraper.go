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

	"github.com/gocolly/colly"
)

/*
func destFile(x string) (f os.File) {
	fName := filepath.Join("D:\\", "go projects", "cwst go", "CWST-GO", "target folder", x)
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, error :%q", err)
	}
	defer file.Close()
	return *file
}

func getSellerName(x string) (y string) {
	var z string
	c := colly.NewCollector(
		colly.AllowedDomains("www.olx.ro"),
	)
	c.OnHTML(".css-1fp4ipz", func(e *colly.HTMLElement) {
		z = e.ChildText("h4")
	})
	c.Visit(x)
	return z
}

func getDateOfPost(x string) (y string) {
	var z string
	c := colly.NewCollector(
		colly.AllowedDomains("www.olx.ro"),
	)
	c.OnHTML(".css-ubdo89-Text", func(e *colly.HTMLElement) {
		z = e.ChildText("span")
	})
	c.Visit(x)
	return z
}

func getItemTitle(x string) (y string) {
	var z string
	c := colly.NewCollector(
		colly.AllowedDomains("www.olx.ro"),
	)
	c.OnHTML(".css-sg1fy9", func(e *colly.HTMLElement) {
		z = e.ChildText("h1")
	})
	c.Visit(x)
	return z
}

func getItemPrice(x string) (y string) {
	var z string
	c := colly.NewCollector(
		colly.AllowedDomains("www.olx.ro"),
	)
	c.OnHTML(".css-dcwlyx", func(e *colly.HTMLElement) {
		z = e.ChildText("h3")
	})
	c.Visit(x)
	return z
}

func getItemDescription(x string) (y string) {
	var z string
	c := colly.NewCollector(
		colly.AllowedDomains("www.olx.ro"),
	)
	c.OnHTML(".css-1m8mzwg", func(e *colly.HTMLElement) {
		z = e.ChildText("div")
	})
	c.Visit(x)
	return z
}
*/

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
	fName := filepath.Join("D:\\", "go projects", "cwst go", "CWST-GO", "target folder", "index.csv")
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, error :%q", err)
	}
	defer file.Close()

	//writer that writes the collected data into our file
	writer := csv.NewWriter(file)

	//channel for output(seller name)
	//y := make(chan string)

	//after the file is written, what it is in the buffer goes in writer and then passed to file
	defer writer.Flush()

	//collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.olx.ro"),
	)

	//HTML parser for seller name
	c.OnHTML(".css-1fp4ipz", func(e *colly.HTMLElement) { //class that contains wanted info

		//y <- e.ChildText("h4")

		//z := <-y

		writer.Write([]string{
			e.ChildText("h4"), //specific tag of the info

			//z,  instead of previous line(not working though)

			//e.ChildText("span"),
		})
	})

	//HTML parser for the date on which the item was posted online
	c.OnHTML(".css-ubdo89-Text", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("span"), //specific tag of the info
		})
	})

	//HTML parser for item title
	c.OnHTML(".css-sg1fy9", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("h1"), //specific tag of the info
		})
	})

	//HTML parser for item price
	c.OnHTML(".css-dcwlyx", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("h3"), //specific tag of the info
		})
	})

	//HTML parser for item description
	c.OnHTML(".css-1m8mzwg", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.ChildText("div"), //specific tag of the info
		})
	})

	//HTML parser for location(not working ¯\_(ツ)_/¯)
	c.OnHTML(".qa-static-ad-map-container img", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			e.Attr("alt"),
		})
	})

	//base64 encoding for all images(only first accessed????)
	c.OnHTML(".swiper-zoom-container img", func(e *colly.HTMLElement) { //class that contains wanted info
		writer.Write([]string{
			imgtoBase64(e.Attr("src")),
		})
	})

	//visiting 3 target pages
	fmt.Printf("Scraping page 1 ... \n")
	c.Visit("https://www.olx.ro/d/oferta/bmw-m-750-xi-garantie-5-ani-full-option-44000km-istoric-service-bmw-IDgqWYY.html")
	/*fmt.Printf("Scraping page 2 ... \n")
	c.Visit("https://www.olx.ro/d/oferta/televizor-smart-qled-samsung-75qn900a-189-cm-ultra-hd-8k-neo-qled-IDfR6sI.html")
	fmt.Printf("Scraping page 3 ... \n")
	c.Visit("https://www.olx.ro/d/oferta/suport-tac-biliard-din-2-bucati-IDgo4s9.html")*/

	log.Printf("\n\nScraping Complete\n\n")
	log.Println(c)

}
