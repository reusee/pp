package pp

// Src provides stream of values
//
// returning value may be nil, which does not indicate the end of stream
// returning nil Src indicates the end of stream
type Src[T any] func() (*T, Src[T], error)

// Get returns the next non-null value, or returns nil if Src is nil
func Get[T any](src *Src[T]) (value *T, err error) {
	for value == nil {
		if src != nil && *src != nil {
			value, *src, err = (*src)()
		} else {
			break
		}
	}
	return
}
