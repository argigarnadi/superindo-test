package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"superindo-test/model"
	"superindo-test/service"
)

type CartController interface {
	AddCart(c *gin.Context)
	GetCart(c *gin.Context)
}

type cartController struct {
	cartServ service.CartService
}

func NewCartController(cartServ service.CartService) CartController {
	return &cartController{
		cartServ: cartServ,
	}
}

func (cc *cartController) AddCart(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if userId == nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "user id not found",
		})
		return
	}

	var req model.RequestCart
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "invalid request",
		})
		return
	}

	req.UserId = userId.(string)
	data, err := cc.cartServ.AddCart(req)
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

func (cc *cartController) GetCart(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if userId == nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "user id not found",
		})
		return
	}

	data, err := cc.cartServ.GetCart(userId.(string))
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
