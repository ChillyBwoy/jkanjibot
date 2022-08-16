package models

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
