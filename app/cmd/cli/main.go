package main

import (
	"go-concurrent-importer/config"
	"go-concurrent-importer/container"
	"go-concurrent-importer/internal/cli"
	"log"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	ctn, err := container.New(cfg)
	if err != nil {
		log.Fatal("cannot initiate app: ", err)
	}

	cliHandler := cli.NewCliHandler(ctn, &cfg.CLI)
	cliHandler.Run()
}
