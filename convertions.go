package list

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Select[A, B any](l List[A], mapper func(a A) B) List[B] {
	var result = make(List[B], 0, len(l))

	for _, a := range l {
		result = append(result, mapper(a))
	}

	return result
}

func Max[A constraints.Ordered](l List[A]) (A, error) {
	if len(l) == 0 {
		var a A
		return a, ErrEmptyList
	}

	r := OrderByDesc(l, func(a A) A { return a })

	return r[0], nil
}

func Min[A constraints.Ordered](l List[A]) (A, error) {
	if len(l) == 0 {
		var a A
		return a, ErrEmptyList
	}

	r := OrderByAsc(l, func(a A) A { return a })

	return r[0], nil
}

func MaxByField[A any, B constraints.Ordered](l List[A], get func(a A) B) (A, error) {
	if len(l) == 0 {
		var a A
		return a, ErrEmptyList
	}

	r := OrderByDesc(l, get)

	return r[0], nil
}

func MinByField[A any, B constraints.Ordered](l List[A], get func(a A) B) (A, error) {
	if len(l) == 0 {
		var a A
		return a, ErrEmptyList
	}

	r := OrderByAsc(l, get)

	return r[0], nil
}

func OrderByDesc[A any, B constraints.Ordered](l List[A], get func(a A) B) List[A] {
	sort.Slice(l, func(i, j int) bool {
		return get(l[i]) > get(l[j])
	})

	return l
}

func OrderByAsc[A any, B constraints.Ordered](l List[A], get func(a A) B) List[A] {
	sort.Slice(l, func(i, j int) bool {
		return get(l[i]) < get(l[j])
	})

	return l
}
