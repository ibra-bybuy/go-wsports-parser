package main

import (
	"log"

	"github.com/ibra-bybuy/wsports-parser/internal/controller"
	"github.com/ibra-bybuy/wsports-parser/internal/handler/parser"
	"github.com/ibra-bybuy/wsports-parser/internal/repository/ufc"
)

func main() {
	log.Println("Starting service for fetching events")

	// FETCH MMA
	ufcParser := parser.New()
	ufcRep := ufc.New(ufcParser.Collector)
	mmaController := controller.New(ufcRep)
	mmaEvents := mmaController.GetEvents()
	log.Printf("FETCHED MMA EVENTS %+v\n", mmaEvents)

	// FETCH FOOTBALL
	// footballParser := parser.New()
	// footballRep := football.New(footballParser.Collector)
	// footballController := controller.New(footballRep)
	// footballEvents := footballController.GetEvents()
	// log.Printf("FETCHED FOOTBALL EVENTS %+v\n", footballEvents)

	// ADD EVENTS TO DATABASE

}
