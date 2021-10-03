package pp

func Tap(fn func(any) error) Sink {
	var sink Sink
	sink = func(v any) (Sink, error) {
		if v == nil {
			return nil, nil
		}
		if err := fn(v); err != nil {
			return nil, err
		}
		return sink, nil
	}
	return sink
}
