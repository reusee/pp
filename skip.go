package pp

func SkipSrc[
	T any,
	Src interface {
		~func() (*T, Src, error)
	},
](src Src, n int, cont Src) Src {
	var ret Src
	ret = func() (*T, Src, error) {
		value, err := Get[T, Src](&src)
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
