package transformer

func Slice[In, Out any](items []In, f func(In) Out) []Out {
	result := make([]Out, len(items))
	for i, item := range items {
		result[i] = f(item)
	}

	return result
}
