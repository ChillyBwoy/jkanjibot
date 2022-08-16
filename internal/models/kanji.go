package models

type JLPT int8

const (
	JlptN1 JLPT = iota + 1
	JlptN2
	JlptN3
	JlptN4
	JlptN5
)

type Reading struct {
	Reading string   `json:"reading"`
	Meaning []string `json:"meaning"`
}

type Kanji struct {
	Char    string    `json:"kanji"`
	ExtId   int       `json:"ext_id"`
	Key     int       `json:"key"`
	Strokes int       `json:"strokes"`
	Jlpt    JLPT      `json:"jlpt"`
	Onyomi  []Reading `json:"onyomi"`
	Kunyomi []Reading `json:"kunyomi"`
}
