package list

import "errors"

var (
	ErrNotFound  = errors.New("entity not found")
	ErrEmptyList = errors.New("list is empty")
)

type List[A any] []A

func (l *List[A]) Where(predicate func(a A) bool) *List[A] {
	var result List[A]

	for _, a := range *l {
		if predicate(a) {
			result = append(result, a)
		}
	}

	return &result
}

func (l *List[A]) Any(predicate func(a A) bool) bool {
	for _, a := range *l {
		if predicate(a) {
			return true
		}
	}

	return false
}

func (l *List[A]) All(predicate func(a A) bool) bool {
	for _, a := range *l {
		if !predicate(a) {
			return false
		}
	}

	return true
}

func (l *List[A]) First(predicate func(a A) bool) (A, error) {
	for _, a := range *l {
		if predicate(a) {
			return a, nil
		}
	}

	var r A

	return r, ErrNotFound
}

func (l *List[A]) FirstOrDefault(predicate func(a A) bool) A {
	var r A

	for _, a := range *l {
		if predicate(a) {
			return a
		}
	}

	return r
}
