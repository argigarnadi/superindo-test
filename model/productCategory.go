package model

import "github.com/google/uuid"

type RequestProductCategory struct {
	Name string `json:"name"`
}

type ResponseProductCategory struct {
	ProductCategoryId uuid.UUID `json:"id"`
	Name              string    `json:"name"`
}
