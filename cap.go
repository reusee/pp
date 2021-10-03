package pp

func CapSrc(src Src, n int, cont Src) Src {
	var ret Src
	ret = func() (any, Src, error) {
		if n == 0 {
			return nil, cont, nil
		}
		value, err := src.Next()
		if err != nil {
			return nil, nil, err
		}
		if value == nil && src == nil {
			return nil, cont, nil
		}
		if value != nil {
			n--
		}
		return value, ret, nil
	}
	return ret
}
