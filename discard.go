package pp

func Discard[T any](v *T) (Sink[T], error) {
	if v == nil {
		return nil, nil
	}
	return Discard[T], nil
}

