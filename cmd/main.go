package main

import (
	"log"
	"net/http"
	"parser/internal/business"
	"parser/internal/data"
	"parser/internal/scheduler"
	"parser/internal/service"
)

func main() {
	database := data.NewInMemoryDB()
	ethereum := data.NewEthereum()

	parser := business.NewParser(database, ethereum)

	err := parser.InitCurrentBlockNumber()
	if err != nil {
		panic(err)
	}

	go scheduler.NewScheduler(parser).Run()

	parserService := service.NewParserService(parser)

	routing(parserService)
	log.Printf("server starting on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func routing(parserService *service.ParserService) {
	http.HandleFunc("/current_block", parserService.GetCurrentBlock)
	http.HandleFunc("/subscribe", parserService.Subscribe)
	http.HandleFunc("/transactions", parserService.GetTransactions)
}
