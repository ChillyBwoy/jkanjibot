package main

import (
	"flag"
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
		log.Println("Error loading .env file")
	}
}

func main() {
	var dev bool

	flag.BoolVar(&dev, "dev", false, "dev mode")

	flag.Parse()

	app := server.NewApp(os.Getenv("TELEGRAM_BOT_TOKEN"), true)

	if dev {
		app.Run()
	} else {
		app.HandleWebhook(os.Getenv("TELEGRAM_BOT_HOST"), os.Getenv("TELEGRAM_BOT_PORT"))
	}
}
