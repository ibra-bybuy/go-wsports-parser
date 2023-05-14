package model

type Stream struct {
	Link string `json:"link" bson:"link"`
	Lang Lang   `json:"lang" bson:"lang"`
}
