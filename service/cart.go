package service

import (
	"fmt"
	"superindo-test/model"
	"superindo-test/repositories"
)

type CartService interface {
	AddCart(req model.RequestCart) (res model.ResponseCart, err error)
	GetCart(userId string) (res []model.CartList, err error)
}

type cartService struct {
	cartRepo    repositories.CartRepository
	productRepo repositories.ProductRepository
}

func NewCartService(cartRepo repositories.CartRepository, productRepo repositories.ProductRepository) CartService {
	return &cartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (cs *cartService) AddCart(req model.RequestCart) (res model.ResponseCart, err error) {
	cart, err := cs.cartRepo.AddCart(req)
	if err != nil {
		fmt.Printf("======> error message : %v", err)
		return
	}

	product, err := cs.productRepo.GetProductById(req.ProductId)

	res = model.ResponseCart{
		Product:  product.Name,
		Quantity: cart.Quantity,
		CreateAt: product.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return
}

func (cs *cartService) GetCart(userId string) (res []model.CartList, err error) {
	res, err = cs.cartRepo.GetListCart(userId)
	if err != nil {
		fmt.Printf("======> error message : %v", err)
		return
	}
	return
}
