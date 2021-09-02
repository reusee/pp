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

func CountSrc[
	T any,
	Src interface {
		~func() (*T, Src, error)
	},
](n *int, src Src, cont Src) Src {
	*n = 0
	var ret Src
	ret = func() (*T, Src, error) {
		value, err := Get[T, Src](&src)
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
