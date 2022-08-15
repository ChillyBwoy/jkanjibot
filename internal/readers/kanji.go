package readers

import (
	"encoding/json"
	"jkanjibot/internal/models"
	"os"
)

func ReadKanji() (*models.Kanji, error) {
	file, err := os.ReadFile("data/kanji.json")
	if err != nil {
		return nil, err
	}

	var kanji models.Kanji

	if err := json.Unmarshal([]byte(file), &kanji.Items); err != nil {
		return nil, err
	}

	return &kanji, nil
}
