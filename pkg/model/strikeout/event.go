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

const HideElements = ".navbar:0,.row:0,.row:1,.mt-1:0,.row:3,.row:4,.row:5,.col-lg-3:0,.btn-danger:0,.d-none:0"

func (e *Event) ToEvent(allItems *[]model.Event, lang model.Lang) (model.Event, error) {
	filterName := strings.ReplaceAll(e.Name, ".", "")
	filterName = strings.ReplaceAll(filterName, " W ", " ")
	names := strings.Split(filterName, "vs")
	eStartTime, _ := e.GetTime()

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
				itemTime, _ := item.GetTime()
				sameHour := datetime.SameDayHour(eStartTime, itemTime)
				if strings.Contains(team1Name, strings.ToLower(name1)) || strings.Contains(team2Name, strings.ToLower(name2)) && sameHour {

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
	t, err := datetime.FromYMDHS(e.StartAtDateTime)
	newTime := t.Add(-time.Hour)
	return newTime, err
}
