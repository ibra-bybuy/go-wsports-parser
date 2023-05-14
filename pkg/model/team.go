package model

type Team struct {
	ID        string `json:"id" bson:"id"`
	Name      string `json:"name" bson:"name"`
	AvatarURL string `json:"avatarUrl" bson:"avatarUrl"`
	Lang      Lang   `json:"lang" bson:"lang"`
}
