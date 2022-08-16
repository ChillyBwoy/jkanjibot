package containers

import "math/rand"

type ItemContainer[T any] struct {
	Items *[]T
}

func (k *ItemContainer[T]) RandomSet(n int) []T {
	size := len(*k.Items)
	idx := rand.Perm(size)[:n]

	pick := make([]T, 0, n)

	for _, idx := range idx {
		pick = append(pick, (*k.Items)[idx])
	}

	return pick
}
