package pp

// Get returns the next non-null value, or returns nil if Src is nil
func Get[
	T any,
	Src interface {
		~func() (*T, Src, error)
	},
](src *Src) (value *T, err error) {
	for value == nil {
		if src != nil && *src != nil {
			value, *src, err = (*src)()
		} else {
			break
		}
	}
	return
}
