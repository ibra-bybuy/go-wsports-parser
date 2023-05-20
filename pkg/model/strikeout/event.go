package strikeout

import (
	"errors"
	"strings"
	"time"

	"github.com/ibra-bybuy/wsports-parser/pkg/model"
	"github.com/ibra-bybuy/wsports-parser/pkg/utils/datetime"
)

type Event struct {
	URL             string `json:"url"`
	Name            string `json:"name"`
	StartAtDateTime string `json:"startAtDateTime"`
	StartAtTime     string `json:"startAtTime"`
}

type Events []Event

const HideElements = ".navbar, .row.text-center.mt-2, .row .col-12 h2, .row .col-12.text-center, .row .col-lg-3, .row .col-lg-9 .mt-1"

func (e *Event) ToEvent(allItems *[]model.Event, lang model.Lang) (model.Event, error) {
	filterName := strings.ReplaceAll(e.Name, ".", "")
	filterName = strings.ReplaceAll(filterName, " W ", " ")
	names := strings.Split(filterName, "vs")

	if len(names) >= 2 {
		name1 := strings.TrimSpace(names[0])
		name2 := strings.ReplaceAll(strings.TrimSpace(names[1]), " W", "")

		returnItem := model.Event{
			StartAt: e.StartAtDateTime,
			Teams: []model.Team{
				{
					ID:   name1,
					Name: name1,
					Lang: lang,
				},
				{
					ID:   name2,
					Name: name2,
					Lang: lang,
				},
			},
		}
		for _, item := range *allItems {
			if len(item.Teams) >= 2 {
				team1Name := strings.ToLower(item.Teams[0].Name)
				team2Name := strings.ToLower(item.Teams[1].Name)
				if strings.Contains(team1Name, strings.ToLower(name1)) || strings.Contains(team2Name, strings.ToLower(name2)) {
					item.HideElements = HideElements
					returnItem = item
				}
			}
		}

		returnItem.Streams = []model.Stream{
			{
				Link: e.URL,
				Lang: lang,
			},
		}

		if returnItem.Name != "" {
			return returnItem, nil
		}
	}

	return model.Event{}, errors.New("Error")
}

func (evs *Events) ToEvents(allItems *[]model.Event, lang model.Lang) *[]model.Event {
	events := []model.Event{}

	for _, event := range *evs {
		newEvents, err := event.ToEvent(allItems, lang)
		if err == nil {
			events = append(events, newEvents)
		}

	}

	return &events
}

func (evs *Events) GetByTerms(terms []string) *Event {
	var event Event
	lastItemContainingWords := 0

	for _, e := range *evs {
		containingWords := 0
		for _, term := range terms {
			if strings.Contains(strings.ToLower(e.Name), strings.ToLower(term)) {
				containingWords += 1
			}
		}

		if containingWords > lastItemContainingWords {
			lastItemContainingWords = containingWords
			event = e
		}
	}

	if lastItemContainingWords < 1 {
		return nil
	}

	return &event
}

func (e *Event) GetTime() (time.Time, error) {
	return datetime.FromYMDHS(e.StartAtDateTime)
}
