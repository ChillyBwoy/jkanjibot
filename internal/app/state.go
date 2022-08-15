package app

import "jkanjibot/internal/quiz"

type AppState struct {
	HiraganaQuiz *quiz.KanaQuiz
	KanjiQuiz    *quiz.KanjiQuiz
}
