package solid

import (
	"fmt"
)

type Vector3 struct {
	X, Y, Z float32
}

func (v Vector3) String() string {
	return fmt.Sprintf("%#v", v)
}
