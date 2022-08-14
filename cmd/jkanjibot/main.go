package main

import (
	"jkanjibot/internal/app"
	"jkanjibot/internal/commands"
	"jkanjibot/internal/quiz"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (
	appState *app.AppState
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appState = &app.AppState{
		HiraganaQuiz: quiz.NewHKanaQuiz("data/hiragana.json"),
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	commandsMap, commandsList := commands.MakeCommands()

	setCommands := tgbotapi.NewSetMyCommands(commandsList...)

	if _, err := bot.Request(setCommands); err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.IsCommand() {
			cmd := update.Message.Command()

			matchedCommand := commandsMap[cmd]
			if matchedCommand != nil {
				if err := matchedCommand.Handler(appState, bot, &update); err != nil {
					log.Panic(err)
				}
			}
		}
	}
}