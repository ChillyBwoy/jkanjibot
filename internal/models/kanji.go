package models

import (
	"math/rand"
)

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

type KanjiChar struct {
	Char    string    `json:"kanji"`
	ExtId   int       `json:"ext_id"`
	Key     int       `json:"key"`
	Strokes int       `json:"strokes"`
	Jlpt    JLPT      `json:"jlpt"`
	Onyomi  []Reading `json:"onyomi"`
	Kunyomi []Reading `json:"kunyomi"`
}

type Kanji struct {
	Items *[]KanjiChar
}

func (k *Kanji) RandomSet(n int) []KanjiChar {
	size := len(*k.Items)
	idx := rand.Perm(size)[:n]

	pick := make([]KanjiChar, 0, n)

	for _, idx := range idx {
		pick = append(pick, (*k.Items)[idx])
	}

	return pick
}
