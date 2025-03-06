package model

type RequestCart struct {
	UserId    string `json:"userId"`
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type ResponseCart struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
	CreateAt string `json:"createAt"`
}

type CartList struct {
	ProductId       string  `json:"productId" gorm:"column:product_id"`
	ProductName     string  `json:"productName" gorm:"column:name"`
	ProductQuantity int     `json:"quantity" gorm:"column:quantity"`
	ProductPrice    float64 `json:"price" gorm:"column:price"`
	TotalPrice      float64 `json:"totalPrice"`
}
