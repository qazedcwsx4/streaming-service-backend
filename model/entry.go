package model

type Entry struct {
	EntryId     string `db:"entry.id" json:"entry_id"`
	Name        string `db:"name" json:"name"`
	ImageUrl    string `db:"image" json:"image_url"`
	Description string `db:"description" json:"description"`
	episodes    []Episode
}
