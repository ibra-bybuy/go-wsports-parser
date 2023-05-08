package goalcom

import (
	model "github.com/ibra-bybuy/wsports-parser/pkg/model/goalcom"
)

type repository struct {
	date   string
	offset int32
}

func New(date string) *repository {
	return &repository{
		date:   date,
		offset: 180,
	}
}

func (r *repository) Get() *model.LivescoreList {
	response := model.Response{}
	r.request(&response)
	return (*model.LivescoreList)(&response.Livescores)
}
