package pp

func Tap(fn func(any)) Sink {
	var sink Sink
	sink = func(v any) (Sink, error) {
		if v == nil {
			return nil, nil
		}
		fn(v)
		return sink, nil
	}
	return sink
}
