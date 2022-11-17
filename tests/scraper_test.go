package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDestFile(t *testing.T) {
	/*if destFile("index.csv") == nil {
		t.Error("destination file could not be created successfully")
	}*/
	assert.Equal(t, destFile("index.csv"), "index3.yaml")
}

func TestGetSellerName(t *testing.T) {
	/*if getSellerName("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html") == nil {
		t.Error("could not get seller name")
	}*/
	assert.Equal(t, getSellerName("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html"), "avenue")
}

func TestGetDateOfPost(t *testing.T) {
	/*if getDateOfPost("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html") == nil {
		t.Error("could not get the date on which the item was posted online")
	}*/
	assert.Equal(t, getDateOfPost("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html"), "Valabil din 18.10.2022 15:28:31")
}

func TestGetItemTitle(t *testing.T) {
	/*if getItemTitle("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html") == nil {
		t.Error("could not get item price")
	}*/
	assert.Equal(t, getItemTitle("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html"), "Grapa cu discuri Grano System de la 2.7m")
}

func TestGetItemPrice(t *testing.T) {
	/*if getItemPrice("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html") == nil {
		t.Error("could not get item price")
	}*/
	assert.Equal(t, getItemPrice("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html"), "6000EUR")
}

func TestGetItemDescription(t *testing.T) {
	/*if getItemDescription("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html") == nil {
		t.Error("could not get item description")
	}*/
	assert.Equal(t, getItemDescription("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html"), "AlteleUtilaje agricole www.bestauto.ro MarcaAlteleModelgrapa cu discuri grano system de la 2.7m Categoria Utilaje agricole Sub categoria alte masini agricole Avariatneavariata  Lucrarile agricole specifice toamnei presupun eliberarea terenurilor de resturile vegetale si efectuarea araturilor cu incorporarea reziduurilor in sol pentru mineralizare optima. Astfel va prezentam o gama diversa de grape cu discuri cultivatoare combinatoare scarificatoare de calitate superioara fabricate in Polonia Dotate cu -discuri OFAS diametre disponibile 510 560 È™i 610 mm din oÈ›el special tratat pentru a conferi rezistenta crescuta- butuci prevazuti cu rulmenti capsati lubrifiati si sigilati in baie de ulei pentru cresterea durabilitatii si a calitatii economisirea timpului si minimalizarea costurilor de intretinere-protectie de cauciuc impotriva suprasarcinii si a socurilor-ecrane lateral de protectie-constructie solida-pliere hidraulica pentru facilitarea transportului-sistem hidraulic pentru atasarea semanatoarei (optional)-cÄƒrucior de transport cu barÄƒ de tracÈ›iune Precum si inovatii Recipient pentru raspandirea ingrasamintelor atasat pe utilaj cu o capacitate a rezervorului de 110 litri lÄƒÈ›ime de rÄƒspÃ¢ndire 2 2m-15m Toate acestea menite sa simplifice si sa modernizeze viata agricultorului.Discul se utilizeazÄƒ si la alte lucrÄƒri la spargerea crustei formate la suprafa a solului grÄƒparea pÄƒ unilor grÄƒparea semÄƒnÄƒturilor si a fÃ¢ne elor naturale precum si Ã®ngroparea Ã®ngrÄƒ Äƒmintelor si eventual a semintelor distribuite pe suprafa a solului. Materialele utilizate sunt de foarta buna calitate cea ce ofera grapei cu disc o fiabilitate foarte buna.         FACILITATI Posibilitate achizitie prin factura externa (TVA 0%) sau prin Fonduri Europene.   Se ofera Certificat conformitate si calitate de la producator.   Livrare in toate regiunile tarii la preturi rezonabile.   Posibilitate achizitie prin rate TBI Publi24_1659956376Publi24_165995637607xx xxx xxxArata telefon")
}

func TestGetItemLocation(t *testing.T) {
	/*if getItemLocation("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html") == nil {
		t.Error("could not get item price")
	}*/
	assert.Equal(t, getItemLocation("https://www.publi24.ro/anunturi/auto-moto/utilaje/utilaje-agricole/alte-masini-agricole/anunt/grapa-cu-discuri-grano-system-de-la-2-7m/id59h447906g7e8gd91e64e53922i087.html"), "Caras-SeverinCaransebes")
}
