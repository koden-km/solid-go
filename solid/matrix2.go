package solid

import (
	"fmt"
)

type Matrix2 [4]float32

func (m Matrix2) String() string {
	return fmt.Sprintf("%#v", m)
}
