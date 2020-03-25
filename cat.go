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
