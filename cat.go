package pp

func CatSrc(srcs ...Src) Src {
	var fn Src
	fn = func() (any, Src, error) {
		if len(srcs) == 0 {
			return nil, nil, nil
		}
		if srcs[0] == nil {
			srcs = srcs[1:]
			return nil, fn, nil
		}
		var value any
		var err error
		value, srcs[0], err = srcs[0]()
		if err != nil {
			return nil, nil, err
		}
		return value, fn, nil
	}
	return fn
}

func CatSink(sinks ...Sink) Sink {
	var ret Sink
	ret = func(value any) (Sink, error) {
		if len(sinks) == 0 {
			return nil, nil
		}
		if value == nil {
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
