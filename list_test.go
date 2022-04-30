package list

import (
	"testing"
)

func TestWhere(t *testing.T) {
	l := NewList[int]()
	l = append(l, 1, 2, 3, 4, 5, 6)

	wl := l.Where(func(a int) bool {
		if a > 5 {
			return true
		}
		return false
	})

	if len(wl) != 1 {
		t.Errorf("len of list more then 1 and equal %v", len(wl))
	}

	if wl[0] != 6 {
		t.Errorf("fist value in list is %v and it is not equal to 6", wl[0])
	}
}

func TestAny(t *testing.T) {
	l := NewList[int]()
	l = append(l, 1, 2, 3, 4, 5, 6)

	result := l.Any(func(a int) bool {
		if a == 5 {
			return true
		}
		return false
	})

	if !result {
		t.Error("result is not true")
	}
}

func TestAll(t *testing.T) {
	l := NewList[int]()
	l = append(l, 1, 2, 3, 4, 5, 6)

	result := l.All(func(a int) bool {
		if a != 10 {
			return true
		}
		return false
	})

	if !result {
		t.Error("result is not true")
	}
}

func TestFirst(t *testing.T) {
	l := NewList[int]()
	l = append(l, 1, 2, 3, 4, 5, 6)

	result, err := l.First(func(a int) bool {
		if a == 3 {
			return true
		}
		return false
	})

	if err != nil {
		t.Error("an error has occurred")
	}

	if result != 3 {
		t.Error("result is not 3")
	}
}

func TestFirstOrDefault(t *testing.T) {
	l := NewList[int]()
	l = append(l, 1, 2, 3, 4, 5, 6)

	result := l.FirstOrDefault(func(a int) bool {
		if a == 10 {
			return true
		}
		return false
	})

	if result != 0 {
		t.Error("result is not 0")
	}
}
