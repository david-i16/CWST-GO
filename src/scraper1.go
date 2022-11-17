package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"image"
	"image/jpeg"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var scrapedPageNum int
var item_folder_name string
var item_folder_name1 string
var imageNum int

/*func getItemLocation(website string) string {
	var x string
	c := colly.NewCollector(
		colly.AllowedDomains("www.publi24.ro"),
	)
	c.OnHTML(".medium-5", func(e *colly.HTMLElement) { //class that contains wanted info
		x = e.ChildText("a")
	})
	c.Visit("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html")
	return x
}*/

/*func createDestFileAndWriter(name string) csv.Writer {
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
}*/

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

func modelate_item_folder_name(x string) string {
	t := strings.Replace(x, "Valabil din ", "", -1)
	z := strings.ReplaceAll(t, " ", "_")
	return z
}

func modelate_item_folder_name1(x string) string {
	z := strings.ReplaceAll(x, ".", "_")
	y := strings.ReplaceAll(z, ":", "_")
	return y
}

func base64toJpg(data string) {

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println("base64toJpg", bounds, formatString)

	//Encode from image format to writer
	imageNum++
	pngFilename := "image_" + strconv.Itoa(imageNum) + ".jpg"

	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Jpg file", pngFilename, "created")
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
		item_folder_name1 = "-" + e.ChildText("i")
		writer.Write([]string{
			e.ChildText("i"), //specific tag of the info
		})
	})

	//HTML parser for item title
	c.OnHTML(".large-8", func(e *colly.HTMLElement) { //class that contains wanted info
		item_folder_name = e.ChildText("h1")
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
		base64String := imgtoBase64(e.Attr("src"))
		base64toJpg(base64String[strings.IndexByte(base64String, ',')+1:])
		writer.Write([]string{
			base64String,
		})
	})

	//visiting 3 target pages
	fmt.Printf("Scraping page 1 ... \n")
	c.Visit("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html")
	log.Printf("\n\nScraping Complete\n\n")
	log.Println(c)

	item_folder_name = modelate_item_folder_name(item_folder_name + modelate_item_folder_name1(item_folder_name1))
	if err := os.MkdirAll("TARGET FOLDER/"+item_folder_name, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	//scrape("https://www.publi24.ro/anunturi/auto-moto/masini-second-hand/vw/golf/anunt/vw-golf-6-diesel-euro-5/8248h8271g4875fhd1d60789036fe1h7.html")
}
