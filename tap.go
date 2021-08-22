package pp

func Tap[T any](fn func(*T) error) Sink[T] {
	var sink Sink[T]
	sink = func(v *T) (Sink[T], error) {
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
