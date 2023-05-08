package model

type Team struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
	Lang      Lang   `json:"lang"`
}
