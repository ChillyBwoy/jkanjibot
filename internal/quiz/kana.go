package quiz

import (
	"jkanjibot/internal/containers"
	"jkanjibot/internal/models"
	"jkanjibot/internal/readers"
	"log"
	"math/rand"
)

type KanaQuiz struct {
	Hiragana *containers.ItemContainer[models.Mora]
}

func NewKanaQuiz(path string) *KanaQuiz {
	reader := &readers.CharReader[models.Mora]{}

	hiragana, err := reader.Read(path)

	if err != nil {
		log.Panic(err)
	}

	cmd := &KanaQuiz{
		Hiragana: hiragana,
	}

	return cmd
}

func (c *KanaQuiz) GetPayload() (*models.QuizQuestion, error) {
	moras := c.Hiragana.RandomSet(4)

	choices := make([]string, 0)

	for _, m := range moras {
		choices = append(choices, m.Romaji)
	}

	idx := rand.Intn(len(moras))

	question := &models.QuizQuestion{
		Question: moras[idx].Kana,
		Choices:  choices,
		Answer:   idx,
	}

	return question, nil
}
