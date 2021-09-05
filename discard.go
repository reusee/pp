package pp

func Discard[
	T any,
	Sink interface {
		~func(*T) (Sink, error)
	},
](v *T) (Sink, error) {
	if v == nil {
		return nil, nil
	}
	return Discard[T, Sink], nil
}

