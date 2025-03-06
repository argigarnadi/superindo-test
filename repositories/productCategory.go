package repositories

import (
	"gorm.io/gorm"
	"log"
	"superindo-test/model"
	"time"
)

type ProductCategoryRepository interface {
	AddCategory(req model.RequestProductCategory) (data model.ResponseProductCategory, err error)
	GetAllCategory() (data []model.ProductCategory, err error)
	GetCategoryById(id string) (data model.ProductCategory, err error)
}

type productCategoryRepository struct {
	DB *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) ProductCategoryRepository {
	return &productCategoryRepository{
		DB: db,
	}
}

func (pcr *productCategoryRepository) AddCategory(req model.RequestProductCategory) (data model.ResponseProductCategory, err error) {
	productCategory := model.ProductCategory{
		Name:      req.Name,
		CreatedAt: time.Now(),
	}

	err = pcr.DB.Create(&productCategory).Error
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}

	data = model.ResponseProductCategory{
		ProductCategoryId: productCategory.ProductCategoryId,
		Name:              productCategory.Name,
	}
	return
}

func (pcr *productCategoryRepository) GetAllCategory() (data []model.ProductCategory, err error) {
	if err = pcr.DB.Find(&data).Error; err != nil {
		log.Printf("======> error message : %v", err)
		return
	}
	return
}

func (pcr *productCategoryRepository) GetCategoryById(id string) (data model.ProductCategory, err error) {
	if err = pcr.DB.Where("product_category_id = ?", id).First(&data).Error; err != nil {
		log.Printf("======> error message : %v", err)
		return
	}
	return
}
