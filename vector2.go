package solid

import (
	"fmt"
)

type Vector2 struct {
	X, Y float32
}

func (v Vector2) String() string {
	return fmt.Sprintf("v2{%#v,%#v}", v.X, v.Y)
}
