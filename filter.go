package pp

func FilterSrc[
	T any,
	Src interface {
		~func() (*T, Src, error)
	},
](
	src Src,
	predict func(T) bool,
	cont Src,
) Src {
	var ret Src
	ret = func() (*T, Src, error) {
		value, err := Get[T, Src](&src)
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

func FilterSink[
	T any,
	Sink interface {
		~func(*T) (Sink, error)
	},
](
	sink Sink,
	predict func(T) bool,
) Sink {
	var ret Sink
	ret = func(value *T) (Sink, error) {
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
