package list

func Select[A, B any](l List[A], mapper func(a A) B) List[B] {
	var result List[B]

	for _, a := range l {
		result = append(result, mapper(a))
	}

	return result
}
