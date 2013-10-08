package solid

import (
	"fmt"
)

type Quaternion struct {
	X, Y, Z, W float32
}

func (q Quaternion) String() string {
	return fmt.Sprintf("q{%#v,%#v,%#v,%#v}", q.X, q.Y, q.Z, q.W)
}
