package pp

func MapSrc(
	src Src,
	fn func(any) any,
	cont Src,
) Src {
	var ret Src
	ret = func() (any, Src, error) {
		value, err := src.Next()
		if err != nil {
			return nil, nil, err
		}
		if value == nil && src == nil {
			return nil, cont, nil
		}
		if value != nil {
			value = fn(value)
		}
		return value, ret, nil
	}
	return ret
}
