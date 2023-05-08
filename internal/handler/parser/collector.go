package parser

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Handler struct {
	*colly.Collector
}

func New() *Handler {
	return &Handler{setupCollector()}
}

func setupCollector() *colly.Collector {
	c := colly.NewCollector()
	c.AllowURLRevisit = true

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response from", r.Request.URL, "status code", r.StatusCode)
	})

	return c
}
