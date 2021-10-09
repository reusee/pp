package pp

func CapSrc[
	Src interface {
		~func() (*T, Src, error)
	},
	T any,
](src Src, n int, cont Src) Src {
	var ret Src
	ret = func() (*T, Src, error) {
		if n == 0 {
			return nil, cont, nil
		}
		value, err := Get[Src](&src)
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
