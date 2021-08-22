package pp

func CatSrc[T any](srcs ...Src[T]) Src[T] {
	var fn Src[T]
	fn = func() (*T, Src[T], error) {
		if len(srcs) == 0 {
			return nil, nil, nil
		}
		if srcs[0] == nil {
			srcs = srcs[1:]
			return nil, fn, nil
		}
		var value *T
		var err error
		value, srcs[0], err = srcs[0]()
		if err != nil {
			return nil, nil, err
		}
		return value, fn, nil
	}
	return fn
}

func CatSink[T any](sinks ...Sink[T]) Sink[T] {
	var ret Sink[T]
	ret = func(value *T) (Sink[T], error) {
		if value != nil && len(sinks) == 0 {
			return nil, ErrShortSink
		}
		if len(sinks) == 0 {
			return nil, nil
		}
		var err error
		sinks[0], err = sinks[0](value)
		if err != nil {
			return nil, err
		}
		if sinks[0] == nil {
			sinks = sinks[1:]
		}
		return ret, nil
	}
	return ret
}
