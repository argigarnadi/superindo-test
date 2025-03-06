package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"superindo-test/model"
	"superindo-test/service"
)

type ProductController interface {
	Create(c *gin.Context)
	GetProduct(c *gin.Context)
	GetProductById(c *gin.Context)
}

type productController struct {
	productServ service.ProductService
}

func NewProductController(productServ service.ProductService) ProductController {
	return &productController{productServ: productServ}
}

func (p *productController) Create(c *gin.Context) {
	isAdmin, exists := c.Get("is_admin")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if !isAdmin.(bool) {
		c.JSON(http.StatusUnauthorized, model.FailedResponse{
			Status:  "failed",
			Message: "you are not admin",
		})
		return
	}

	var req model.RequestProduct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "invalid request",
		})
		return
	}

	data, err := p.productServ.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedResponse{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{
		Status: "success",
		Result: data,
	})
	return
}

func (p *productController) GetProduct(c *gin.Context) {
	data, err := p.productServ.GetProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedResponse{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{
		Status: "success",
		Result: data,
	})
}

func (p *productController) GetProductById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "id must be provided",
		})
		return
	}

	data, err := p.productServ.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedResponse{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{
		Status: "success",
		Result: data,
	})
}
