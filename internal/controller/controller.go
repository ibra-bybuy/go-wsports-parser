package controller

import (
	"github.com/ibra-bybuy/wsports-parser/pkg/model"
)

type repository interface {
	Get() (*[]model.Event, error)
}

type Controller struct {
	r repository
}

func New(r repository) *Controller {
	return &Controller{r}
}

func (c *Controller) GetEvents() *[]model.Event {
	items, err := c.r.Get()

	if err != nil {
		return &[]model.Event{}
	}

	return items
}
