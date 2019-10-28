package main

import "puzzles-golang/patterns/value-receivers/content"

func main() {

	// The fact that we use a value receiver will cause the value of l to be copied, so technically, what we set pageSize to is a completely different place in memory. This is why we have to return a pointer to it and and reassign l.
	l := content.NewLoader().WithPageSize(5)

	_ = l.Load([]byte("some bytes"))
}

// This will cause some copying and a bit of work for the garbage collector, but in the grand scheme of things, it will be a negligible overhead. The safety that you gain is far more important.
