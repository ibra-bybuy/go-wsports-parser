package main

import (
	"context"
	"log"
	"time"

	fetchController "github.com/ibra-bybuy/wsports-parser/internal/controller/fetch_controller"
	saveController "github.com/ibra-bybuy/wsports-parser/internal/controller/save_controller"
	"github.com/ibra-bybuy/wsports-parser/internal/handler/parser"
	"github.com/ibra-bybuy/wsports-parser/internal/repository/dotenv"
	"github.com/ibra-bybuy/wsports-parser/internal/repository/events"
	"github.com/ibra-bybuy/wsports-parser/internal/repository/football"
	"github.com/ibra-bybuy/wsports-parser/internal/repository/mma"
	"github.com/ibra-bybuy/wsports-parser/internal/repository/mongodb"
	"github.com/ibra-bybuy/wsports-parser/pkg/model"
)

func main() {
	log.Println("Starting service for fetching events")

	// Load .env vars
	dotenv.Load()

	// INITING DATABASE
	mngClient := mongodb.New()
	sRep := events.NewMongo(mngClient)
	sCtrl := saveController.New(sRep)

	defer func() {
		if err := mngClient.Disconnect(mngClient.Ctx); err != nil {
			panic(err)
		}
	}()

	// Fetching data every 6 hours
	for {
		allEvents := []model.Event{}
		// FETCH MMA
		mmaParser := parser.New()
		mmaRep := mma.New(mmaParser.Collector)
		mmaController := fetchController.New(mmaRep)
		mmaEvents := mmaController.GetEvents()
		allEvents = append(allEvents, *mmaEvents...)
		log.Printf("FETCHED MMA EVENTS %+v\n", mmaEvents)

		// FETCH FOOTBALL
		footballParser := parser.New()
		footballRep := football.New(footballParser.Collector)
		footballController := fetchController.New(footballRep)
		footballEvents := footballController.GetEvents()
		allEvents = append(allEvents, *footballEvents...)
		log.Printf("FETCHED FOOTBALL EVENTS %+v\n", footballEvents)

		// ADD EVENTS TO DATABASE
		sCtrl.Add(context.Background(), &allEvents)
		time.Sleep(time.Hour * 6)
	}

}
