package readers

import (
	"encoding/json"
	"jkanjibot/internal/models"
	"os"
)

func ReadKana(path string) (*models.Kana, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var kana models.Kana

	if err := json.Unmarshal([]byte(file), &kana.Chars); err != nil {
		return nil, err
	}

	return &kana, nil
}
