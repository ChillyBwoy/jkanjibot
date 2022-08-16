package server

import (
	"encoding/json"
	"fmt"
	"io"
	"jkanjibot/internal/app"
	"jkanjibot/internal/commands"
	"jkanjibot/internal/quiz"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct {
	Bot         *tgbotapi.BotAPI
	state       *app.AppState
	commandsMap map[string]commands.Command
}

func NewApp(token string, debug bool) *App {
	state := &app.AppState{
		HiraganaQuiz: quiz.NewKanaQuiz("data/hiragana.json"),
		KanjiQuiz:    quiz.NewKanjiQuiz("data/kanji.json"),
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	commandsMap, commandsList := commands.MakeCommands()
	setCommands := tgbotapi.NewSetMyCommands(commandsList...)

	if _, err := bot.Request(setCommands); err != nil {
		log.Fatal(err)
	}

	bot.Debug = debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &App{
		commandsMap: commandsMap,
		Bot:         bot,
		state:       state,
	}
}

func (a *App) dispatch(update tgbotapi.Update) *tgbotapi.Message {
	if update.Message == nil {
		return nil
	}

	if !update.Message.IsCommand() {
		return nil
	}

	cmd := update.Message.Command()
	matchedCommand := a.commandsMap[cmd]

	if matchedCommand == nil {
		return nil
	}

	msg, err := matchedCommand.Handler(a.state, a.Bot, &update)
	if err != nil {
		log.Fatal("Unable to handle command", err)
	}

	return msg
}

func (a *App) Handle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)

	var update tgbotapi.Update
	if err := json.Unmarshal(body, &update); err != nil {
		log.Fatal("Error en el update", err)
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := a.dispatch(update)
	if msg != nil {
		w.Header().Add("Content-Type", "application/json")
		log.Printf("%+v\n", msg)
	}
}

func (a *App) RunServer(host, port string) {
	addr := fmt.Sprintf("%s:%s/%s", host, port, a.Bot.Token)
	wh, _ := tgbotapi.NewWebhook(addr)

	_, err := a.Bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := a.Bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := a.Bot.ListenForWebhook("/" + a.Bot.Token)
	go http.ListenAndServe("0.0.0.0:8443", nil)

	for update := range updates {
		a.dispatch(update)
	}
}

func (a *App) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := a.Bot.GetUpdatesChan(u)

	for update := range updates {
		a.dispatch(update)
	}
}
