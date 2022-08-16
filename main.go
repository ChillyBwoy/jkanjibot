package main

import (
	"jkanjibot/internal/server"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := server.NewApp(os.Getenv("TELEGRAM_BOT_TOKEN"), false)
	app.RunServer(os.Getenv("TELEGRAM_BOT_DOMAIN"), os.Getenv("TELEGRAM_BOT_PORT"))
}
