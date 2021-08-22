package pp

func MapSrc[T any](
	src Src[T],
	fn func(T) T,
	cont Src[T],
) Src[T] {
	var ret Src[T]
	ret = func() (*T, Src[T], error) {
		value, err := src.Next()
		if err != nil {
			return nil, nil, err
		}
		if value == nil && src == nil {
			return nil, cont, nil
		}
		if value != nil {
			*value = fn(*value)
		}
		return value, ret, nil
	}
	return ret
}

func MapSink[T any](
	sink Sink[T],
	fn func(T) T,
) Sink[T] {
	var ret Sink[T]
	ret = func(value *T) (Sink[T], error) {
		if value != nil && sink == nil {
			return nil, ErrShortSink
		}
		var err error
		if value != nil {
			sink, err = sink(PtrOf(fn(*value)))
		} else {
			sink, err = sink(nil)
		}
		if err != nil {
			return nil, err
		}
		if sink == nil {
			return nil, nil
		}
		return ret, nil
	}
	return ret
}
