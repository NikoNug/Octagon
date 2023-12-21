package models

type History struct {
	ID       int `json:"ID"`
	PostID   int `json:"post_id"`
	PersonID int `json:"person_id"`
}
