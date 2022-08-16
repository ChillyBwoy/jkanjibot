package commands

import (
	"jkanjibot/internal/app"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command interface {
	GetDescription() string
	GetCommand() string
	BotCommand() tgbotapi.BotCommand
	Handler(appState *app.AppState, bot *tgbotapi.BotAPI, update *tgbotapi.Update) (*tgbotapi.Message, error)
}

type CommandMap = map[string]Command
type CommandList = []tgbotapi.BotCommand

func MakeCommands() (CommandMap, CommandList) {
	hiraganaCommand := NewHiraganaCommand()
	kanjiCommand := NewKanjiCommand()

	commandsMap := make(map[string]Command)
	commandsMap[hiraganaCommand.GetCommand()] = hiraganaCommand
	commandsMap[kanjiCommand.GetCommand()] = kanjiCommand

	commandsList := make([]tgbotapi.BotCommand, 0)

	for _, cmd := range commandsMap {
		commandsList = append(commandsList, cmd.BotCommand())
	}

	return commandsMap, commandsList
}
