package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"superindo-test/model"
	"time"
)

type ProductRepository interface {
	AddProduct(req model.RequestProduct) (data model.ResponseProduct, err error)
	GetAllProduct() (data []model.Product, err error)
	GetProductById(id string) (data model.Product, err error)
}

type productRepository struct {
	DB               *gorm.DB
	prodCategoryRepo ProductCategoryRepository
}

func NewProductRepository(db *gorm.DB, prodCategoryRepoo ProductCategoryRepository) ProductRepository {
	return &productRepository{
		DB:               db,
		prodCategoryRepo: prodCategoryRepoo,
	}
}

func (pr *productRepository) AddProduct(req model.RequestProduct) (data model.ResponseProduct, err error) {

	prodcutCategory, err := uuid.Parse(req.ProductCategoryId)
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}

	product := model.Product{
		Name:              req.Name,
		ImageUrl:          req.ImageUrl,
		Price:             req.Price,
		ProductCategoryId: prodcutCategory,
		CreatedAt:         time.Now(),
	}

	err = pr.DB.Create(&product).Error
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}

	getCategory, err := pr.prodCategoryRepo.GetCategoryById(req.ProductCategoryId)
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}

	data = model.ResponseProduct{
		ProductId:         product.ProductId,
		Name:              product.Name,
		ImageUrl:          product.ImageUrl,
		Price:             product.Price,
		ProductCategoryId: getCategory.Name,
		CreateAt:          product.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return
}

func (pr *productRepository) GetAllProduct() (data []model.Product, err error) {
	if err = pr.DB.Find(&data).Error; err != nil {
		log.Printf("======> error message : %v", err)
		return
	}
	return
}

func (pr *productRepository) GetProductById(id string) (data model.Product, err error) {
	if err = pr.DB.Where("product_id = ?", id).First(&data).Error; err != nil {
		log.Printf("======> error message : %v", err)
		return
	}
	return
}
