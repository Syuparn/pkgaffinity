package baz

// break rule
import (
	"a/foo/bar"
	"a/other/hoge"
)

func NewBar() *bar.Bar {
	var _ = hoge.Hoge
	return &bar.Bar{}
}
