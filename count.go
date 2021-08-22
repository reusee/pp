package pp

func CountSink[T any](n *int) Sink[T] {
	*n = 0
	var ret Sink[T]
	ret = func(value *T) (Sink[T], error) {
		if value != nil {
			*n++
			return ret, nil
		} else {
			return nil, nil
		}
	}
	return ret
}

func CountSrc[T any](n *int, src Src[T], cont Src[T]) Src[T] {
	*n = 0
	var ret Src[T]
	ret = func() (*T, Src[T], error) {
		value, err := src.Next()
		if err != nil {
			return nil, nil, err
		}
		if value != nil {
			*n++
		} else if src == nil {
			return nil, cont, nil
		}
		return value, ret, nil
	}
	return ret
}
