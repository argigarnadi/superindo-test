package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"superindo-test/model"
)

func LoadConfig() (config *viper.Viper, err error) {
	config = viper.New()
	config.SetConfigFile("app.yaml")
	config.SetConfigType("yaml")
	config.AutomaticEnv()

	err = config.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file, %s", err)
	}
	log.Println("Configuration loaded successfully")
	return
}

func InitDatabase(viper *viper.Viper) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetString("database.port"))

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	fmt.Println("Connected to Database")
	// Create the extension for UUID generation if it doesn't exist
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	fmt.Println("Successfully create extension")

	// Migrate the schema
	err = db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.ProductCategory{},
		&model.Cart{},
		&model.ProductDetail{},
	)
	if err != nil {
		log.Fatal("failed migration schema")
	}
	fmt.Println("Migration schema successfully")
	return
}

func InitServer() *gin.Engine {
	server := gin.Default()
	return server
}
