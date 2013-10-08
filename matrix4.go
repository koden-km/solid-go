package solid

import (
	"fmt"
)

type Matrix4 [16]float32

func (m Matrix4) String() string {
	return fmt.Sprintf("%#v", m)
}
