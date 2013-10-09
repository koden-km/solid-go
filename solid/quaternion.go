package solid

import (
	"fmt"
)

type Quaternion struct {
	X, Y, Z, W float32
}

func (q Quaternion) String() string {
	return fmt.Sprintf("%#v", q)
}
