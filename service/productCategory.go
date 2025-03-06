package service

import (
	"log"
	"superindo-test/model"
	"superindo-test/repositories"
)

type ProductCategoryService interface {
	Create(req model.RequestProductCategory) (data model.ResponseProductCategory, err error)
	GetCategory() (res []model.ResponseProductCategory, err error)
}

type productCategoryService struct {
	prodCatRepo repositories.ProductCategoryRepository
}

func NewProductCategoryService(prodCatRepo repositories.ProductCategoryRepository) ProductCategoryService {
	return &productCategoryService{
		prodCatRepo: prodCatRepo,
	}
}

func (pc *productCategoryService) Create(req model.RequestProductCategory) (data model.ResponseProductCategory, err error) {
	data, err = pc.prodCatRepo.AddCategory(req)
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}
	return
}

func (pc *productCategoryService) GetCategory() (res []model.ResponseProductCategory, err error) {
	data, err := pc.prodCatRepo.GetAllCategory()
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}
	for _, v := range data {
		res = append(res, model.ResponseProductCategory{
			ProductCategoryId: v.ProductCategoryId,
			Name:              v.Name,
		})
	}
	return
}
