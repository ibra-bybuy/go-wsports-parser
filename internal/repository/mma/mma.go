package mma

import (
	"strings"

	"github.com/gocolly/colly"
	"github.com/ibra-bybuy/wsports-parser/internal/repository/strikeout"
	"github.com/ibra-bybuy/wsports-parser/internal/repository/ufc"
	"github.com/ibra-bybuy/wsports-parser/pkg/model"
	strikeoutModel "github.com/ibra-bybuy/wsports-parser/pkg/model/strikeout"
	"github.com/ibra-bybuy/wsports-parser/pkg/utils/datetime"
)

type repository struct {
	c *colly.Collector
}

func New(c *colly.Collector) *repository {
	return &repository{c}
}

func (r *repository) Get() (*[]model.Event, error) {
	ufcEvents := r.getUFCEvents()
	strikeoutEvents := r.getStrikeoutEvents(*ufcEvents)

	return strikeoutEvents, nil
}

func (r *repository) getUFCEvents() *[]model.Event {
	events, _ := ufc.New(r.c).Get()
	return events
}

func (r *repository) getStrikeoutEvents(allItems []model.Event) *[]model.Event {
	request := strikeout.New(r.c, "https://strikeout.ws/ufc")
	strikeoutEvents := strikeoutModel.Events{}
	request.Request(&strikeoutEvents)

	if len(strikeoutEvents) > 0 {
		for i, event := range allItems {

			foundItem := strikeoutEvents.GetByTerms(strings.Split(event.Name, " "))
			if foundItem != nil {
				t, err := event.GetTime()

				if err != nil || datetime.IsTodayOrTomorrow(t) == false {
					continue
				}

				allItems[i].HideElements = strikeoutModel.HideElements
				allItems[i].Streams = append(event.Streams, model.Stream{
					Lang: model.LangEng,
					Link: foundItem.URL,
				})
			}
		}
	}

	return &allItems
}
