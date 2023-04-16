package main

import (
	"MyGram/app"
	"MyGram/config"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	err := config.InitPostgres()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
