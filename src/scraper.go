package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gocolly/colly"
)

/*func destFile(x string) (f os.File){
	fName := filepath.Join("D:\\", "go projects", "cwst go", "CWST-GO", "target folder", x)
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, error :%q", err)
	}
	defer file.Close()
	return *file
}*/

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
	//after the file is written, what it is in the buffer goes in writer and then passed to file
	defer writer.Flush()

	//collector
	c := colly.NewCollector(
		colly.AllowedDomains("https://www.olx.ro/"),
	)

	//HTML parser
	c.OnHTML(".css-1fp4ipz", func(e *colly.HTMLElement) { //div class that contains wanted info

		writer.Write([]string{
			e.ChildText("h4"), //specific tag of the info
			//e.ChildText("span"),
		})
	})

	fmt.Printf("Scraping page :  ")
	c.Visit("https://www.olx.ro/d/oferta/bmw-xdrixe-seria-7-2020-71000-tva-IDgp7iN.html")

	log.Printf("\n\nScraping Complete\n\n")
	log.Println(c)

}
