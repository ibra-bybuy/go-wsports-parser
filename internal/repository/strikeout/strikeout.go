package strikeout

import (
	"github.com/gocolly/colly"
	"github.com/ibra-bybuy/wsports-parser/pkg/model/strikeout"
)

const BASE_URL = "https://strikeout.ws"

type repository struct {
	c   *colly.Collector
	URL string
}

func New(c *colly.Collector, url string) *repository {
	return &repository{c, url}
}

func (r *repository) Request(events *strikeout.Events) {

	r.c.OnRequest(func(r *colly.Request) {})

	r.c.OnHTML("a.mb-1.btn", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		if url != "" {
			url = BASE_URL + url
		}
		name := e.Attr("title")
		dateTime := e.ChildAttr("span[content]", "content")
		time := e.ChildText("span[content]")

		item := strikeout.Event{
			URL:             url,
			Name:            name,
			StartAtDateTime: dateTime,
			StartAtTime:     time,
		}
		*events = append(*events, item)
	})

	r.c.Visit(r.URL)
}
