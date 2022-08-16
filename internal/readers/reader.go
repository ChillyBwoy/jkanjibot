package readers

import (
	"encoding/json"
	"jkanjibot/internal/containers"
	"os"
)

type CharReader[T any] struct {
}

func (r *CharReader[T]) Read(path string) (*containers.ItemContainer[T], error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var items containers.ItemContainer[T]

	if err := json.Unmarshal([]byte(file), &items.Items); err != nil {
		return nil, err
	}

	return &items, nil
}
