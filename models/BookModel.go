package models

import (
	"time"
)

type Book struct {
	ID 			string `gorm:"primaryKey" json:"id"`
	BookTitle   string `json:"book_title"`
	BookCover 	string `json:"book_cover"`
	Description string `json:"description"`
	AuthorID    uint   `json:"author_id"`
	Author Author `gorm:"foreignKey:AuthorID" json:"author"`
	TotalPages	int 	`json:"total_pages"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

type BookResponse struct {
	ID 			string `json:"id"`
	BookTitle   string `json:"book_title"`
	BookCover 	string `json:"book_cover"`
	Description string `json:"description"`
	AuthorID    uint   `json:"-"`
	Author 		AuthorBookResponse `gorm:"foreignKey:AuthorID" json:"author"`
	TotalPages	int 	`json:"total_pages"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}