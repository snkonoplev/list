package list

import (
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	l := List[int]{1, 2, 3, 4, 5, 6}

	result := Select(l, func(a int) string {
		return fmt.Sprint(a)
	})

	if len(result) != len(l) {
		t.Error("lenhts of lists are not equal")
	}
}
