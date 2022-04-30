package list

func Select[A, B any](l List[A], mapper func(a A) B) List[B] {
	var result = make(List[B], 0, len(l))

	for _, a := range l {
		result = append(result, mapper(a))
	}

	return result
}
