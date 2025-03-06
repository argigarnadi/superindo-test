package model

import "github.com/google/uuid"

type RequestProduct struct {
	Name              string  `json:"name"`
	ImageUrl          string  `json:"imageUrl"`
	Price             float64 `json:"price"`
	ProductCategoryId string  `json:"productCategoryId"`
}

type ResponseProduct struct {
	ProductId         uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	ImageUrl          string    `json:"imageUrl"`
	Price             float64   `json:"price"`
	ProductCategoryId string    `json:"category"`
	CreateAt          string    `json:"createAt"`
}
