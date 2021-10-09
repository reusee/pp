package pp

func Discard[
	Sink interface {
		~func(*T) (Sink, error)
	},
	T any,
](v *T) (Sink, error) {
	if v == nil {
		return nil, nil
	}
	return Discard[Sink], nil
}
