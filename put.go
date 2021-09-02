package pp

// Put puts an element to Sink
func Put[
	T any,
	Sink interface {
		~func(*T) (Sink, error)
	},
](sink *Sink, value *T) (err error) {
	if sink == nil || *sink == nil {
		return nil
	}
	*sink, err = (*sink)(value)
	return
}
