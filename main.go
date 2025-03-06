package main

import (
	"fmt"
	"log"
	"superindo-test/config"
)

func main() {
	//init load config
	loadConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	//init database connection
	db, err := config.InitDatabase(loadConfig)
	if err != nil {
		log.Fatal(err)
	}
	server := config.InitServer()

	config.App(&config.AppConfig{
		Server:     server,
		DB:         db,
		LoadConfig: loadConfig,
	})

	server.Run(fmt.Sprintf(":%d", loadConfig.GetInt("app.port")))

}
