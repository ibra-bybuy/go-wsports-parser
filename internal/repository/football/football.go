package football

import (
	"time"

	"github.com/gocolly/colly"
	"github.com/ibra-bybuy/wsports-parser/internal/repository/goalcom"
	"github.com/ibra-bybuy/wsports-parser/internal/repository/strikeout"
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
	allEvents := getAllEvents()
	strikeoutEvents := r.getStrikeoutEvents(allEvents)

	return strikeoutEvents, nil
}

func getAllEvents() *[]model.Event {
	todayEvents := goalcom.New(datetime.YearMonthDay(time.Now())).Get()
	tomorrowEvents := goalcom.New(datetime.YearMonthDay(time.Now().Add(time.Hour * 24))).Get()
	allEvents := append(*todayEvents, *tomorrowEvents...)
	return allEvents.ToEvents(model.LangEng)
}

func (r *repository) getStrikeoutEvents(allItems *[]model.Event) *[]model.Event {
	footballRequest := strikeout.New(r.c, "https://strikeout.ws/soccer")
	strikeoutEvents := strikeoutModel.Events{}
	footballRequest.Request(&strikeoutEvents)

	return strikeoutEvents.ToEvents(allItems, model.LangEng)
}
