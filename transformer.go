package transformer

func Slice[In, Out any](items []In, f func(In) Out) []Out {
	result := make([]Out, len(items), len(items))
	for key, value := range items {
		result[key] = f(value)
	}

	return result
}
