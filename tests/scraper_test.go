package main

import "testing"

func TestDestFile(t *testing.T) {
	if destFile("index.csv") == nil {
		t.Error("destination file could not be created successfully")
	}
}
