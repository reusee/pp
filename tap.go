package pp

func Tap[
	T any,
	Sink interface {
		~func(*T) (Sink, error)
	},
](fn func(*T) error) Sink {
	var sink Sink
	sink = func(v *T) (Sink, error) {
		if v == nil {
			return nil, nil
		}
		if err := fn(v); err != nil {
			return nil, err
		}
		return sink, nil
	}
	return sink
}
