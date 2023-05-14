package saveController

import (
	"context"

	"github.com/ibra-bybuy/wsports-parser/pkg/model"
)

type repository interface {
	Add(ctx context.Context, events *[]model.Event) bool
}

type Controller struct {
	r repository
}

func New(r repository) *Controller {
	return &Controller{r}
}

func (c *Controller) Add(ctx context.Context, events *[]model.Event) bool {
	return c.r.Add(ctx, events)
}
