package models

import "time"

type App struct {
	ID        int       `json:"id"`
	IdHash    string    `json:"idHash"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
