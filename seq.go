package pp

func Seq[T any](values ...T) Src[T] {
	var src Src[T]
	src = func() (*T, Src[T], error) {
		if len(values) == 0 {
			return nil, nil, nil
		}
		value := values[0]
		values = values[1:]
		return &value, src, nil
	}
	return src
}
