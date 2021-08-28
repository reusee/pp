package pp

func SkipSrc[T any](src Src[T], n int, cont Src[T]) Src[T] {
	var ret Src[T]
	ret = func() (*T, Src[T], error) {
		value, err := Get(&src)
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
