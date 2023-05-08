package repository

import "github.com/ibra-bybuy/wsports-parser/pkg/model"

type Event interface {
	Get() ([]model.Event, error)
}
