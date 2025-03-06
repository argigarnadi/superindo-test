package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"superindo-test/model"
	"time"
)

type UserRepository interface {
	Register(req model.RequestRegister) (data model.ResponseRegister, err error)
	Login(req model.RequestLogin) (data model.User, err error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) Register(req model.RequestRegister) (data model.ResponseRegister, err error) {
	user := model.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}
	err = ur.DB.Create(&user).Error
	if err != nil {
		fmt.Sprintf("======> error message : %v", err)
		return
	}

	data = model.ResponseRegister{
		UserId:   user.UserId,
		Name:     user.Name,
		Email:    user.Email,
		CreateAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return
}

func (ur *userRepository) Login(req model.RequestLogin) (data model.User, err error) {
	if err = ur.DB.Where("email = ? AND password = ?", req.Email, req.Password).First(&data).Error; err != nil {
		err = errors.New("email and password doesn't match")
		return
	}

	return
}
