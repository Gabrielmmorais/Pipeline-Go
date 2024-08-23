package scraper

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func EnviaDados(url string, dataChannel chan<- string) {
	retorno, erro := http.Get(url)
	if erro != nil {
		//fmt.Println("HTML tilte dog", erro)
		dataChannel <- ""
		return
	}
	defer retorno.Body.Close()

	if retorno.StatusCode != 200 {
		//fmt.Println("Deu ruim:", retorno.StatusCode)
		dataChannel <- ""
		return
	}

	dgoquery, erro := goquery.NewDocumentFromReader(retorno.Body)
	if erro != nil {
		//fmt.Println("HTML tilte dog", erro)
		dataChannel <- ""
		return
	}

	texto_titulo := dgoquery.Find("title").Text()

	dataChannel <- texto_titulo
}
