package quiz

import (
	"fmt"
	"jkanjibot/internal/containers"
	"jkanjibot/internal/models"
	"jkanjibot/internal/readers"
	"log"
	"math/rand"
	"strings"
)

type KanjiQuiz struct {
	Kanji *containers.ItemContainer[models.Kanji]
}

func NewKanjiQuiz(path string) *KanjiQuiz {
	reader := &readers.CharReader[models.Kanji]{}

	kanji, err := reader.Read(path)

	if err != nil {
		log.Panic(err)
	}

	cmd := &KanjiQuiz{
		Kanji: kanji,
	}

	return cmd
}

func (c *KanjiQuiz) GetPayload() (*models.QuizQuestion, error) {
	kanji := c.Kanji.RandomSet(4)

	choices := make([]string, 0)

	for _, k := range kanji {
		var onyomi []string
		var kunyomi []string

		for _, reading := range k.Onyomi {
			onyomi = append(onyomi, reading.Reading)
		}

		for _, reading := range k.Kunyomi {
			kunyomi = append(kunyomi, reading.Reading)
		}

		if kunyomi != nil && onyomi != nil {
			choices = append(choices, fmt.Sprintf("%s / %s", strings.Join(onyomi, ", "), strings.Join(kunyomi, ", ")))
		} else if kunyomi != nil {
			choices = append(choices, strings.Join(kunyomi, ", "))
		} else {
			choices = append(choices, strings.Join(onyomi, ", "))
		}
	}

	idx := rand.Intn(len(kanji))

	question := &models.QuizQuestion{
		Question: kanji[idx].Char,
		Choices:  choices,
		Answer:   idx,
	}

	return question, nil
}
