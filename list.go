package list

type List[A any] []A

func NewList[A any]() List[A] {
	l := make(List[A], 0)
	return l
}

func (l *List[A]) Where(predicate func(a A) bool) List[A] {
	var result List[A]

	for _, a := range *l {
		if predicate(a) {
			result = append(result, a)
		}
	}

	return result
}
