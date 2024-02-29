package sliceutil

import (
	"testing"
)

func TestPickRandom(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	random := PickRandom(slice)
	if random == 0 {
		t.Errorf("PickRandom() returned 0")
	}
}
