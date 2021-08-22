package pp

// Sink consumes stream of values
//
// for non-higher-order Sinks, a nil input value indicates the end of stream
// for higher-order Sinks, a nil input value should pass as-is to argument sinks
type Sink[T any] func(*T) (Sink[T], error)

func (s Sink[T]) Sink(v *T) (Sink[T], error) {
	if s == nil {
		return nil, nil
	}
	return s(v)
}

