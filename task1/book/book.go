package book

import "time"

type Book struct {
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Language      string    `json:"language"`
	Publisher     string    `json:"publisher"`
	DatePublished time.Time `json:"date_published"`
}
