package main

import (
	"log"
	"os"
	"path/filepath"
)

func destFile(x string) (f os.File) {
	fName := filepath.Join("D:\\", "go projects", "cwst go", "CWST-GO", "target folder", x)
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, error :%q", err)
	}
	defer file.Close()
	return *file
}

func main() {
	//setting up the file where we store collected data
	//destFile("index.csv")
	//writer that writes the collected data into our file
	//writer := csv.NewWriter(file)
	//after the file is written, what it is in the buffer goes in writer and then passed to file
	//defer writer.Flush()
}
