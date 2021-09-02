package pp

func MapSrc[
	T any,
	Src interface {
		~func() (*T, Src, error)
	},
](
	src Src,
	fn func(T) T,
	cont Src,
) Src {
	var ret Src
	ret = func() (*T, Src, error) {
		value, err := Get[T, Src](&src)
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

func MapSink[
	T any,
	Sink interface {
		~func(*T) (Sink, error)
	},
](
	sink Sink,
	fn func(T) T,
) Sink {
	var ret Sink
	ret = func(value *T) (Sink, error) {
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
