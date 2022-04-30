package list

import "testing"

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
