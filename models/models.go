package models

type Mood struct {
	ID          int    `json:"id"`
	Mood        string `json:"mood"`
	LastUpdated int    `json:"lastUpdated"`
}

type UpdatedMood struct {
	ID      int    `json:"id"`
	NewMood string `json:"newMood"`
}
