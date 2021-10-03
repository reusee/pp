package pp

func SkipSrc(src Src, n int, cont Src) Src {
	var ret Src
	ret = func() (any, Src, error) {
		value, err := src.Next()
		if err != nil {
			return nil, nil, err
		}
		if value == nil && src == nil {
			return nil, cont, nil
		}
		if n > 0 {
			if value != nil {
				n--
			}
			return nil, ret, nil
		}
		return value, ret, nil
	}
	return ret
}
