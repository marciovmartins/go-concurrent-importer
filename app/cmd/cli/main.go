package main

import (
	"go-concurrent-importer/container"
	"go-concurrent-importer/internal/cli"
	"log"
)

func main() {

	ctn, err := container.NewCliApp()
	if err != nil {
		log.Fatal("nao foi possivel iniciar a app")
	}

	path := "../test/segmentations.csv"
	handler := cli.NewCliHandler(ctn)
	handler.ProcessCSV(path)
}

