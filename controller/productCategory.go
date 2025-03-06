package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"superindo-test/model"
	"superindo-test/service"
)

type ProductCategoryController interface {
	Create(c *gin.Context)
	GetCategory(c *gin.Context)
}

type productCategoryController struct {
	prodCateService service.ProductCategoryService
}

func NewProductCategoryController(prodCateService service.ProductCategoryService) ProductCategoryController {
	return &productCategoryController{prodCateService: prodCateService}
}

func (p *productCategoryController) Create(c *gin.Context) {

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

	var req model.RequestProductCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "invalid request",
		})
		return
	}

	data, err := p.prodCateService.Create(req)
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

func (p *productCategoryController) GetCategory(c *gin.Context) {
	data, err := p.prodCateService.GetCategory()
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
