package quiz

import (
	"jkanjibot/internal/models"
	"jkanjibot/internal/readers"
	"log"
	"math/rand"
)

type KanaQuiz struct {
	Hiragana *models.Kana
}

func NewHKanaQuiz(path string) *KanaQuiz {
	hiragana, err := readers.ReadKana(path)

	if err != nil {
		log.Panic(err)
	}

	cmd := &KanaQuiz{
		Hiragana: hiragana,
	}

	return cmd
}

func (c *KanaQuiz) GetPayload() (*QuizQuestion, error) {
	moras := c.Hiragana.RandomSet(4)

	choices := make([]string, 0)

	for _, m := range moras {
		choices = append(choices, m.Romaji)
	}

	idx := rand.Intn(len(moras))

	question := &QuizQuestion{
		Question: moras[idx].Kana,
		Choices:  choices,
		Answer:   idx,
	}

	return question, nil
}
