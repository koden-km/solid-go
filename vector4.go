package solid

import (
	"fmt"
)

type Vector4 struct {
	X, Y, Z, W float32
}

func (v Vector4) String() string {
	return fmt.Sprintf("%#v", v)
}
