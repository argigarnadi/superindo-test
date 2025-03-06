package routes

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"superindo-test/controller"
)

type RouterConfig struct {
	Server                    *gin.Engine
	LoadConfig                *viper.Viper
	UserController            controller.UserController
	ProductCategoryController controller.ProductCategoryController
	ProductController         controller.ProductController
	CartController            controller.CartController
}

func (rc *RouterConfig) Route() {
	rc.Server.POST("/register", rc.UserController.UserRegister)
	rc.Server.POST("/login", rc.UserController.Login)

	gRoute := rc.Server.Group("/")
	gRoute.Use(rc.AuthMiddleware())
	{
		// product category
		gRoute.POST("/create/product-category", rc.ProductCategoryController.Create)
		gRoute.GET("/product-category", rc.ProductCategoryController.GetCategory)

		// product
		gRoute.POST("/create/product", rc.ProductController.Create)
		gRoute.GET("/product", rc.ProductController.GetProduct)
		gRoute.GET("/product/detail/:id", rc.ProductController.GetProductById)

		// cart
		gRoute.POST("/cart/add", rc.CartController.AddCart)
		gRoute.GET("/cart/list", rc.CartController.GetCart)
	}
}

func (rc *RouterConfig) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secretKey := rc.LoadConfig.GetString("jwt.secret")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthenticated"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			c.Set("is_admin", claims["is_admin"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
		}

		c.Next()
	}
}
