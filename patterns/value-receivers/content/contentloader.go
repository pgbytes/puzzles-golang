package content

type Loader struct {
	pageSize int
}

func NewLoader() *Loader {
	return &Loader{
		pageSize: 10, // default
	}
}

// check the use of a value receiver here
// value receivers have one big advantage, which is that they are safe for concurrent use. What a value receiver would get is a copy of the original value. This makes them perfect for implementing builder methods.
func (l Loader) WithPageSize(ps int) *Loader {
	l.pageSize = ps
	return &l
}

// The fact that we use a value receiver will cause the value of l to be copied, so technically, what we set pageSize to is a completely different place in memory. This is why we have to return a pointer to it and and reassign l
// check the main calling method for sample.

// The rest could be your usual pointer receivers
func (l *Loader) Load(bytes []byte) string {
	// ...
	return string(bytes)
}
