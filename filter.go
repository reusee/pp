package pp

func FilterSrc[T any](
	src Src[T],
	predict func(T) bool,
	cont Src[T],
) Src[T] {
	var ret Src[T]
	ret = func() (*T, Src[T], error) {
		value, err := src.Next()
		if err != nil {
			return nil, nil, err
		}
		if src == nil && value == nil {
			return nil, cont, nil
		}
		if value != nil && !predict(*value) {
			value = nil
		}
		return value, ret, nil
	}
	return ret
}

func FilterSink[T any](
	sink Sink[T],
	predict func(T) bool,
) Sink[T] {
	var ret Sink[T]
	ret = func(value *T) (Sink[T], error) {
		if value != nil && sink == nil {
			return nil, ErrShortSink
		}
		var err error
		if value == nil || predict(*value) {
			if sink == nil {
				return nil, nil
			}
			sink, err = sink(value)
			if err != nil {
				return nil, err
			}
		}
		if sink == nil {
			return nil, nil
		}
		return ret, nil
	}
	return ret
}
