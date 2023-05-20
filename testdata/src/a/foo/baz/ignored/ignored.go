package ignored

// break rule but ignored
import (
	"a/foo/bar"
)

func NewBar() *bar.Bar {
	return &bar.Bar{}
}
