package models

import (
	"math/rand"
)

type MoraKind string

const (
	G MoraKind = "gojūon"
	D MoraKind = "dakuon"
	H MoraKind = "handakuon"
	S MoraKind = "sokuon"
	Y MoraKind = "yōon"
)

type Mora struct {
	Kana   string   `json:"kana"`
	Kind   MoraKind `json:"type"`
	Romaji string   `json:"romaji"`
}

type Kana struct {
	Chars *[]Mora
}

func (kana *Kana) RandomMora() *Mora {
	size := len(*kana.Chars)
	idx := rand.Intn(size)

	pick := (*kana.Chars)[idx]

	return &pick
}

func (kana *Kana) RandomSet(n int) []Mora {
	size := len(*kana.Chars)
	idxs := rand.Perm(size)[:n]

	pick := make([]Mora, 0, n)

	for _, idx := range idxs {
		pick = append(pick, (*kana.Chars)[idx])
	}

	return pick
}
