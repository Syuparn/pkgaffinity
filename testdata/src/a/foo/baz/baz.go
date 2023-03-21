package baz

// break rule
import "a/foo/bar"

func NewBar() *bar.Bar {
	return &bar.Bar{}
}
