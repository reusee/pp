package pp

func Discard(v any) (Sink, error) {
	if v == nil {
		return nil, nil
	}
	return Discard, nil
}
