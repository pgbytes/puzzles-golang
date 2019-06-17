//+build mage
// This go file is ignored by the go build build with the above comment.

package main

import "github.com/magefile/mage/sh"

// Runs go build

func Test() error {
	return sh.Run("go", "test", "./...")
}

func Bench() error {
	return sh.Run("go", "test", "./...", "-bench=.")
}
