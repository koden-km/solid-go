package solid

import (
	"fmt"
)

type Matrix3 [9]float32

func (m Matrix3) String() string {
	return fmt.Sprintf("%#v", m)
}
