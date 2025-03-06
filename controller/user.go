package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"superindo-test/constant"
	"superindo-test/model"
	"superindo-test/service"
)

type UserController interface {
	UserRegister(c *gin.Context)
	Login(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (uc *userController) UserRegister(c *gin.Context) {
	var req model.RequestRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "invalid request",
		})
		return
	}

	data, err := uc.userService.Register(req)
	if err != nil {
		fmt.Errorf("error message : %v", err)
		if strings.Contains(err.Error(), constant.DuplicateEmailSQL) {
			c.JSON(http.StatusInternalServerError, model.FailedResponse{
				Status:  "failed",
				Message: "Email already registered",
			})
			return
		}
		c.JSON(http.StatusBadRequest, model.FailedResponse{
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

func (uc *userController) Login(c *gin.Context) {
	var req model.RequestLogin
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "Invalid input request",
		})
		return
	}

	data, err := uc.userService.UserLogin(req)
	if err != nil {
		fmt.Errorf("login failed, err : %v", err)
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
