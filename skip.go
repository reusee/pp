package pp

func SkipSrc[
	Src interface {
		~func() (*T, Src, error)
	},
	T any,
](src Src, n int, cont Src) Src {
	var ret Src
	ret = func() (*T, Src, error) {
		value, err := Get[Src](&src)
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
