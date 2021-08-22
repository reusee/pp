package pp

func CapSrc[T any](src Src[T], n int, cont Src[T]) Src[T] {
	var ret Src[T]
	ret = func() (*T, Src[T], error) {
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
