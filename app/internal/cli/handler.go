package cli

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"go-concurrent-importer/config"
	"go-concurrent-importer/container"
	"go-concurrent-importer/internal/service"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type cliHandler struct {
	segmentationService *service.SegmentationService
	cfg                 *config.CLI
}

func NewCliHandler(ctn *container.Container, cfg *config.CLI) *cliHandler {
	return &cliHandler{
		ctn.SegmentationService,
		cfg,
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
func (ch *cliHandler) Run() {
	fmt.Printf("Bem vindo ao %s versão: %s \n", ch.cfg.Name, ch.cfg.TagVersion)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Informe o caminho do arquivo CSV: ")
		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)

		info, err := os.Stat(path)
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				fmt.Println("Arquivo não encontrado. Tente novamente.")
				continue
			}
			fmt.Printf("Erro ao acessar arquivo: %v\n", err)
			continue
		}

		if info.IsDir() {
			fmt.Println("O caminho informado é um diretório, não um arquivo.")
			continue
		}

		fmt.Println("Arquivo encontrado, aguarde a importação")
		ch.processCSV(path)
	}
}

func (ch *cliHandler) processCSV(path string) {
	start := time.Now()

	// Channels
	recordsChan := make(chan []string, ch.cfg.NumWorkers*2)
	errorsChan := make(chan []error, 100)

	// Init Workers
	var wg sync.WaitGroup
	for i := 0; i < ch.cfg.NumWorkers; i++ {
		wg.Add(1)
		go ch.worker(recordsChan, errorsChan, &wg)
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

	elapsed := time.Since(start)
	log.Printf("tempo decorrido %s", elapsed)
}

func (ch *cliHandler) readCSVStreaming(path string, out chan<- []string) error {
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

func (ch *cliHandler) collectErrors(errs <-chan []error) {
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
