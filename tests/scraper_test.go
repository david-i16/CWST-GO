package main

import "testing"

func TestDestFile(t *testing.T) {
	if destFile("index.csv") == nil {
		t.Error("destination file could not be created successfully")
	}
}

func TestGetSellerName(t *testing.T) {
	if getSellerName("https://www.olx.ro/d/oferta/bmw-xdrixe-seria-7-2020-71000-tva-IDgp7iN.html") == nil {
		t.Error("could not get seller name")
	}
}

func TestGetDateOfPost(t *testing.T) {
	if getDateOfPost("https://www.olx.ro/d/oferta/bmw-xdrixe-seria-7-2020-71000-tva-IDgp7iN.html") == nil {
		t.Error("could not get the date on which the item was posted online")
	}
}

func TestGetItemTitle(t *testing.T) {
	if getItemTitle("https://www.olx.ro/d/oferta/bmw-xdrixe-seria-7-2020-71000-tva-IDgp7iN.html") == nil {
		t.Error("could not get item title")
	}
}
