package pp

// Sink consumes stream of values
//
// for non-higher-order Sinks, a nil input value indicates the end of stream
// for higher-order Sinks, a nil input value should pass as-is to argument sinks
type Sink[T any] func(*T) (Sink[T], error)

// Put puts an element to Sink
func Put[T any](sink *Sink[T], value *T) (err error) {
  if sink == nil || *sink == nil {
    return nil
  }
  *sink, err = (*sink)(value)
  return
}

