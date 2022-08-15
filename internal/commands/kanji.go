package commands

import (
	"jkanjibot/internal/app"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type KanjiCommand struct {
	Command     string
	Description string
}

func NewKanjiCommand() *KanjiCommand {
	return &KanjiCommand{
		Command:     "kanji",
		Description: "Show kanji",
	}
}

func (c *KanjiCommand) GetCommand() string {
	return c.Command
}

func (c *KanjiCommand) GetDescription() string {
	return c.Description
}

func (c *KanjiCommand) BotCommand() tgbotapi.BotCommand {
	return tgbotapi.BotCommand{
		Command:     c.Command,
		Description: c.Description,
	}
}

func (c *KanjiCommand) Handler(appState *app.AppState, bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	question, err := appState.KanjiQuiz.GetPayload()
	if err != nil {
		return err
	}

	poll := tgbotapi.NewPoll(update.Message.Chat.ID, question.Question, question.Choices...)
	poll.Type = "quiz"
	poll.CorrectOptionID = int64(question.Answer)

	if _, err := bot.Send(poll); err != nil {
		return err
	}

	return nil
}
