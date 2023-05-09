package goalcom

import (
	"time"

	"github.com/ibra-bybuy/wsports-parser/pkg/model"
	"github.com/ibra-bybuy/wsports-parser/pkg/utils/datetime"
	"github.com/ibra-bybuy/wsports-parser/pkg/utils/stringformatter"
)

type (
	Response struct {
		Livescores []Livescore `json:"livescores"`
	}

	Livescore struct {
		Competition Competition `json:"competition"`
		Matches     []Match     `json:"matches"`
	}

	LivescoreList []Livescore

	Competition struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Area  Area   `json:"area"`
		Badge Badge  `json:"badge"`
	}

	Area struct {
		Name string `json:"name"`
	}

	Badge struct {
		URL string `json:"url"`
	}

	Match struct {
		ID          string `json:"id"`
		StartDate   string `json:"startDate"`
		Venue       Area   `json:"venue"`
		LastUpdated string `json:"lastUpdated"`
		Round       Round  `json:"round"`
		Status      string `json:"status"`
		TeamA       Team   `json:"teamA"`
		TeamB       Team   `json:"teamB"`
	}

	Round struct {
		Name    string `json:"name"`
		Display bool   `json:"display"`
	}

	Team struct {
		ID    string `json:"id"`
		Code  string `json:"code"`
		Short string `json:"short"`
		Long  string `json:"long"`
		Full  string `json:"full"`
		Name  string `json:"name"`
		Crest Badge  `json:"crest"`
	}
)

func (list *LivescoreList) ToEvents(lang model.Lang) *[]model.Event {
	events := []model.Event{}

	for _, item := range *list {
		for _, match := range item.Matches {
			startAt, err := datetime.FromFull(match.StartDate)

			if err != nil {
				continue
			}

			events = append(events, model.Event{
				ID:        match.TeamA.Full + " - " + match.TeamB.Full,
				Name:      item.Competition.Area.Name + " - " + item.Competition.Name,
				AvatarURL: item.Competition.Badge.URL,
				Teams: []model.Team{
					{
						ID:        match.TeamA.ID,
						Name:      stringformatter.ReplaceI18nLetters(match.TeamA.Full),
						AvatarURL: match.TeamA.Crest.URL,
						Lang:      lang,
					},
					{
						ID:        match.TeamB.ID,
						Name:      stringformatter.ReplaceI18nLetters(match.TeamB.Full),
						AvatarURL: match.TeamB.Crest.URL,
						Lang:      lang,
					},
				},
				StartAt: datetime.Full(startAt),
				EndAt:   datetime.Full(startAt.Add(time.Minute * 150)),
				Address: match.Venue.Name,
				Lang:    lang,
			})
		}
	}

	return &events
}
