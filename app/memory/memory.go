package memory

import (
	"fmt"
)

// ProcessData processa os dados recebidos do scraper
func ProcessData(data string) {
	if data == "" {
		fmt.Println("Dados vazios recebidos, ignorando.")
		return
	}

	// Exemplo de processamento: exibir os dados
	fmt.Println("Processando dados:", data)

	// Aqui você pode fazer qualquer tipo de processamento em memória que precisar
}
