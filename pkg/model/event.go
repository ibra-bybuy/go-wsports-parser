package model

type Event struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Teams     []Team   `json:"teams"`
	StartAt   string   `json:"startAt"`
	EndAt     string   `json:"endAt"`
	AvatarURL string   `json:"avatarUrl"`
	Address   string   `json:"address"`
	Lang      Lang     `json:"lang"`
	Streams   []Stream `json:"stream"`
}
