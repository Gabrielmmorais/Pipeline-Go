package main

import (
	"app/memory"
	"app/scraper"
	"fmt"
)

func main() {
	// URLs para scrapeiar
	urls := []string{
		"https://example.com/page1",
		"https://example.com/page2",
		"https://example.com/page3",
	}

	// Canal para receber os dados scrapeados
	dataChannel := make(chan string)

	// Iniciar goroutine para cada URL
	for _, url := range urls {
		go scraper.EnviaDados(url, dataChannel)
	}

	// Processar os dados conforme são recebidos
	for i := 0; i < len(urls); i++ {
		data := <-dataChannel
		go memory.ProcessData(data)
	}

	// Esperar que todas as goroutines terminem
	fmt.Println("Scraping e processamento finalizados.")
}
