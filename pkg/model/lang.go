package model

type Lang struct {
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
}

var LangRu = Lang{
	Code: "ru",
	Name: "Русский",
}

var LangEng = Lang{
	Code: "en",
	Name: "English",
}
