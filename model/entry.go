package model

type Entry struct {
	EntryId      string `db:"entry.id" json:"entryId"`
	Name         string `db:"name" json:"name"`
	ImageUrl     string `db:"image" json:"imageUrl"`
	Description  string `db:"description" json:"description"`
	EpisodeCount int    `json:"episodeCount"`
	Episodes     []Episode
}
