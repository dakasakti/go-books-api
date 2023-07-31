package models

import "time"

type Author struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthorBookResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}