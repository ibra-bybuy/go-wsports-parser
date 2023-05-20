package model

import (
	"time"

	"github.com/ibra-bybuy/wsports-parser/pkg/utils/datetime"
)

type Event struct {
	ID           string   `json:"id" bson:"id"`
	Name         string   `json:"name" bson:"name"`
	Teams        []Team   `json:"teams" bson:"teams"`
	StartAt      string   `json:"startAt" bson:"startAt"`
	EndAt        string   `json:"endAt" bson:"endAt"`
	AvatarURL    string   `json:"avatarUrl" bson:"avatarUrl"`
	Address      string   `json:"address" bson:"address"`
	Lang         Lang     `json:"lang" bson:"lang"`
	Streams      []Stream `json:"streams" bson:"streams"`
	HideElements string   `json:"hideElements" bson:"hideElements"`
	Sport        string   `json:"sport" bson:"sport"`
}

func (e *Event) GetTime() (time.Time, error) {
	return datetime.FromFull(e.StartAt)
}
