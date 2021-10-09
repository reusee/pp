package pp

func Seq[
	Src interface {
		~func() (*T, Src, error)
	},
	T any,
](values ...T) Src {
	var src Src
	src = func() (*T, Src, error) {
		if len(values) == 0 {
			return nil, nil, nil
		}
		value := values[0]
		values = values[1:]
		return &value, src, nil
	}
	return src
}
