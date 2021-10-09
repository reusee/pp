package pp

func CountSink[
	Sink interface {
		~func(*T) (Sink, error)
	},
	T any,
](n *int) Sink {
	*n = 0
	var ret Sink
	ret = func(value *T) (Sink, error) {
		if value != nil {
			*n++
			return ret, nil
		} else {
			return nil, nil
		}
	}
	return ret
}

func CountSrc[
	Src interface {
		~func() (*T, Src, error)
	},
	T any,
](n *int, src Src, cont Src) Src {
	*n = 0
	var ret Src
	ret = func() (*T, Src, error) {
		value, err := Get[Src](&src)
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
