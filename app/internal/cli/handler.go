package cli

import (
	"bufio"
	"encoding/csv"
	"go-concurrent-importer/container"
	"io"
	"log"
	"os"
	"sync"
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
	// Configs
	numWorkers := 10
	batchSize := 100

	// Channels
	recordsChan := make(chan []string, numWorkers*2)
	errorsChan := make(chan []error, 100)

	// Init Workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go ch.worker(batchSize, recordsChan, errorsChan, &wg)
	}

	// read csv
	readErrChan := make(chan error, 1)
	go func() {
		err := ch.readCSVStreaming(path, recordsChan)
		readErrChan <- err
	}()

	// collects errors
	var errWg sync.WaitGroup
	errWg.Add(1)
	go func() {
		defer errWg.Done()
		ch.collectErrors(errorsChan)
	}()

	// Aguarda leitura terminar e pega o erro
	readErr := <-readErrChan
	if readErr != nil {
		log.Printf("erro ao ler CSV: %v", readErr)
		return
	}

	// Aguarda todos os workers terminarem
	wg.Wait()

	// Fecha channel de erros para sinalizar collectErrors
	close(errorsChan)

	// Aguarda collectErrors terminar
	errWg.Wait()
}

func (ch *cliHandler) readCSVStreaming (path string, out chan<- []string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	defer close(out)

	reader := csv.NewReader(bufio.NewReader(file))
	reader.LazyQuotes = true // Permite aspas não escapadas em campos não-quoted (útil para JSON)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		out <- record
	}

	return nil
}

func (ch *cliHandler) collectErrors (errs <-chan []error){
	var count int

	for err := range errs {
		if err == nil {
			continue
		}
		count++
		log.Printf("erro ao processar registro %v", err)
	}

	if count > 0 {
		log.Printf("processamento concluido com %d erros", count)
	} else {
		log.Printf("processamento concluido sem erros")
	}
}