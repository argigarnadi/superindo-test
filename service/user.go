package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"log"
	"superindo-test/model"
	"superindo-test/repositories"
	"time"
)

type UserService interface {
	Register(req model.RequestRegister) (data model.ResponseRegister, err error)
	UserLogin(req model.RequestLogin) (data model.Token, err error)
}

type userService struct {
	userRepository repositories.UserRepository
	loadConfig     *viper.Viper
}

func NewUserService(uRepo repositories.UserRepository, loadConfig *viper.Viper) UserService {
	return &userService{
		userRepository: uRepo,
		loadConfig:     loadConfig,
	}
}

func (us *userService) Register(req model.RequestRegister) (data model.ResponseRegister, err error) {
	data, err = us.userRepository.Register(req)
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}
	return
}

func (us *userService) UserLogin(req model.RequestLogin) (data model.Token, err error) {
	var user model.User
	user, err = us.userRepository.Login(req)
	if err != nil {
		log.Printf("======> error message : %v", err)
		return
	}

	if user.Email == "" {
		err = errors.New("Email not found")
		log.Printf("======> error message : %v", err)
		return
	}

	data.AccessToken, err = us.GenerateToken(user, time.Hour*1) // Access token valid for 1 hour
	if err != nil {
		err = errors.New("failed generate access token")
		return
	}

	data.RefreshToken, err = us.GenerateToken(user, time.Hour*24*7) // Refresh token valid for 7 days
	if err != nil {
		err = errors.New("failed generate refresh token")
		return
	}

	return
}

func (us *userService) GenerateToken(user model.User, duration time.Duration) (token string, err error) {
	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.UserId,
		"is_admin": user.IsAdmin,
		"exp":      time.Now().Add(duration).Unix(),
	})

	secretKey := us.loadConfig.GetString("jwt.secret")
	token, err = tokenJwt.SignedString([]byte(secretKey))
	if err != nil {
		return
	}

	return
}
