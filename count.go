package pp

func CountSink(n *int) Sink {
	*n = 0
	var ret Sink
	ret = func(value any) (Sink, error) {
		if value != nil {
			*n++
			return ret, nil
		} else {
			return nil, nil
		}
	}
	return ret
}

func CountSrc(n *int, src Src, cont Src) Src {
	*n = 0
	var ret Src
	ret = func() (any, Src, error) {
		value, err := src.Next()
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
