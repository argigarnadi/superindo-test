package service

import (
	"log"
	"superindo-test/model"
	"superindo-test/repositories"
)

type ProductService interface {
	Create(req model.RequestProduct) (data model.ResponseProduct, err error)
	GetProduct() (res []model.ResponseProduct, err error)
	GetProductById(id string) (res model.ResponseProduct, err error)
}

type productService struct {
	productRepo         repositories.ProductRepository
	productCategoryRepo repositories.ProductCategoryRepository
}

func NewProductService(productRepo repositories.ProductRepository, productCateRepo repositories.ProductCategoryRepository) ProductService {
	return &productService{
		productRepo:         productRepo,
		productCategoryRepo: productCateRepo,
	}
}

func (ps *productService) Create(req model.RequestProduct) (data model.ResponseProduct, err error) {
	data, err = ps.productRepo.AddProduct(req)
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}
	return
}

func (ps *productService) GetProduct() (res []model.ResponseProduct, err error) {
	data, err := ps.productRepo.GetAllProduct()
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}

	for _, v := range data {
		category, errCat := ps.productCategoryRepo.GetCategoryById(v.ProductCategoryId.String())
		if errCat != nil {
			log.Printf("======> error message : %v", errCat)
			return
		}

		res = append(res, model.ResponseProduct{
			ProductId:         v.ProductId,
			Name:              v.Name,
			ProductCategoryId: category.Name,
			Price:             v.Price,
			ImageUrl:          v.ImageUrl,
			CreateAt:          v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return
}

func (ps *productService) GetProductById(id string) (res model.ResponseProduct, err error) {
	data, err := ps.productRepo.GetProductById(id)
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}

	category, err := ps.productCategoryRepo.GetCategoryById(data.ProductCategoryId.String())
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}
	res = model.ResponseProduct{
		ProductId:         data.ProductId,
		Name:              data.Name,
		ProductCategoryId: category.Name,
		Price:             data.Price,
		ImageUrl:          data.ImageUrl,
		CreateAt:          data.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return
}
