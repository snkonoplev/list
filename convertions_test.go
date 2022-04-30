package list

import (
	"fmt"
	"testing"
)

type Test struct {
	field int
}

func TestSelect(t *testing.T) {
	l := List[int]{1, 2, 3, 4, 5, 6}

	result := Select(l, func(a int) string {
		return fmt.Sprint(a)
	})

	if len(result) != len(l) {
		t.Error("lenhts of lists are not equal")
	}
}

func TestMax(t *testing.T) {
	l := List[int]{1, 2, 3, 4, 5, 6}

	result, err := Max(l)
	if err != nil {
		t.Error("got error")
	}

	if result != 6 {
		t.Error("max value is incorrect")
	}
}

func TestMin(t *testing.T) {
	l := List[int]{1, 2, 3, 4, 5, 6}

	result, err := Min(l)
	if err != nil {
		t.Error("got error")
	}

	if result != 1 {
		t.Error("min value is incorrect")
	}
}

func TestMaxByField(t *testing.T) {
	l := List[Test]{
		Test{field: 1},
		Test{field: 2},
		Test{field: 3},
	}

	result, err := MaxByField(l, func(a Test) int { return a.field })
	if err != nil {
		t.Error("got error")
	}

	if result.field != 3 {
		t.Error("max value is incorrect")
	}
}

func TestMinByField(t *testing.T) {
	l := List[Test]{
		Test{field: 1},
		Test{field: 2},
		Test{field: 3},
	}

	result, err := MinByField(l, func(a Test) int { return a.field })
	if err != nil {
		t.Error("got error")
	}

	if result.field != 1 {
		t.Error("min value is incorrect")
	}
}
