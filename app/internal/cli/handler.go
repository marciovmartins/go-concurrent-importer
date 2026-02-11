package cli

import (
	"bufio"
	"encoding/csv"
	"go-concurrent-importer/container"
	"log"
	"os"
)

type cliHandler struct {
	cliApp *container.CliApp
}

func NewCliHandler(cliApp *container.CliApp) *cliHandler{
	return &cliHandler{
		cliApp,
	}
}

// ==================== MAIN ====================

// Cenário:
// - Ler um arquivo CSV com 4 colunas: user_id, segment_type, segment_name e data
// - Este arquivo terá 1 milhão de linhas
// - O processamento deve ser performático e otimizado
// - Validar se os dados são válidos
// - Salvar no banco de dados
// - Se houver erro, mostrar ou salvar essa informação de alguma forma
func (ch *cliHandler) ProcessCSV(path string) {
	ctn := ch.cliApp

	arquivo, err := os.Open(path)
	if err != nil {
		log.Printf("open csv: %v", err)
		return
	}
	defer arquivo.Close()

	leitor := csv.NewReader(bufio.NewReader(arquivo))
	records, _ := leitor.ReadAll()

	errs := []error{}

	for _, record := range records {
		err := ctn.SegmentationService.ProcessRecord(record)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		log.Fatalf("Processamento concluído com erros: %v", errs)
	}

	log.Printf("Processamento concluído sem erros")
}