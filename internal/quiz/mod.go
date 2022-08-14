package quiz

type QuizQuestion struct {
	Question string
	Choices  []string
	Answer   int
}

type Quiz interface {
	GetPayload() (*QuizQuestion, error)
}
