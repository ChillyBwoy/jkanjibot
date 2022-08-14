package models

type JLPT int8

const (
	JlptN1 JLPT = iota + 1
	JlptN2
	JlptN3
	JlptN4
	JlptN5
)

type ReadingType string

const (
	Onyomi  ReadingType = "onyomi"
	Kunyomi ReadingType = "kunyomi"
)

type Reading struct {
	Reading string      `json:"reading"`
	Meaning string      `json:"meaning"`
	Kind    ReadingType `json:"kind"`
}

func NewReading(reading, meaning string, kind ReadingType) *Reading {
	return &Reading{
		Reading: reading,
		Meaning: meaning,
		Kind:    kind,
	}
}

type Kanji struct {
	Char     string `json:"char"`
	ExtId    int    `json:"extId"`
	ExtKey   int    `json:"extKey"`
	Strokes  int    `json:"strokes"`
	Jlpt     JLPT   `json:"jlpt"`
	Readings []Reading
}

func NewKanji(char string, ext_id, ext_key, strokes int, jlpt JLPT, readings []Reading) *Kanji {
	return &Kanji{
		Char:     char,
		ExtId:    ext_id,
		ExtKey:   ext_key,
		Strokes:  strokes,
		Jlpt:     jlpt,
		Readings: readings,
	}
}
