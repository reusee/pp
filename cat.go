package pp

func CatSrc[
	Src interface {
		~func() (*T, Src, error)
	},
	T any,
](srcs ...Src) Src {

	if len(srcs) == 0 {
		return nil
	}

	for srcs[0] == nil {
		srcs = srcs[1:]
		if len(srcs) == 0 {
			return nil
		}
	}

	var fn Src
	fn = func() (*T, Src, error) {
		value, err := Get[Src](&srcs[0])
		if err != nil {
			return nil, nil, err
		}
		if value == nil {
			srcs = srcs[1:]
			if len(srcs) == 0 {
				return nil, nil, nil
			}
			return nil, fn, nil
		}
		return value, fn, nil
	}
	return fn
}

func CatSink[
	T any,
	Sink interface {
		~func(*T) (Sink, error)
	},
](sinks ...Sink) Sink {

	if len(sinks) == 0 {
		return nil
	}

	for sinks[0] == nil {
		sinks = sinks[1:]
		if len(sinks) == 0 {
			return nil
		}
	}

	var ret Sink
	ret = func(value *T) (Sink, error) {
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
		for sinks[0] == nil {
			sinks = sinks[1:]
			if len(sinks) == 0 {
				return nil, nil
			}
		}
		return ret, nil
	}
	return ret
}
