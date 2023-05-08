package ufc

import (
	"github.com/gocolly/colly"
	"github.com/ibra-bybuy/wsports-parser/pkg/model"
)

type repository struct {
	c *colly.Collector
}

func New(c *colly.Collector) *repository {
	return &repository{c}
}

func (r *repository) Get() (*[]model.Event, error) {
	events := []model.Event{}
	r.parseEvents(&events, &model.LangEng)

	return &events, nil
}
