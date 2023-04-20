package model

type Book struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type BookStruct struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
