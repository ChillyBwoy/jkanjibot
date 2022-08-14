package commands

import (
	"jkanjibot/internal/app"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HiraganaCommand struct {
	Command     string
	Description string
}

func NewHiraganaCommand() *HiraganaCommand {
	return &HiraganaCommand{
		Command:     "hiragana",
		Description: "Show hiragana",
	}
}

func (c *HiraganaCommand) GetCommand() string {
	return c.Command
}

func (c *HiraganaCommand) GetDescription() string {
	return c.Description
}

func (c *HiraganaCommand) BotCommand() tgbotapi.BotCommand {
	return tgbotapi.BotCommand{
		Command:     c.Command,
		Description: c.Description,
	}
}

func (c *HiraganaCommand) Handler(appState *app.AppState, bot *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	question, err := appState.HiraganaQuiz.GetPayload()
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
