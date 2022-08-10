package main

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

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
}
