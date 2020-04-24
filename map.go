package pp

func MapSrc(
	src Src,
	fn func(any) any,
	cont Src,
) Src {
	var ret Src
	ret = func() (any, Src, error) {
		value, err := src.Next()
		if err != nil {
			return nil, nil, err
		}
		if value == nil && src == nil {
			return nil, cont, nil
		}
		if value != nil {
			value = fn(value)
		}
		return value, ret, nil
	}
	return ret
}

func MapSink(
	sink Sink,
	fn func(any) any,
) Sink {
	var ret Sink
	ret = func(value any) (Sink, error) {
		if value != nil && sink == nil {
			return nil, ErrShortSink
		}
		if value == nil {
			return nil, nil
		}
		var err error
		sink, err = sink(fn(value))
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
