package main

import (
	"TelegramBot/cmd/controller/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	router.Start()
}
