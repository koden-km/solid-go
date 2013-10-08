package solid

import (
	"fmt"
)

type Matrix2 [4]float32

func (m Matrix2) String() string {
	return fmt.Sprintf("m2{{%#v,%#v},{%#v,%#v}}",
		m[0], m[1],
		m[2], m[3])
}
