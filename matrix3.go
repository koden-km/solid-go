package solid

import (
	"fmt"
)

type Matrix3 [9]float32

func (m Matrix3) String() string {
	return fmt.Sprintf("m3{{%#v,%#v,%#v},{%#v,%#v,%#v},{%#v,%#v,%#v}}",
		m[0], m[1], m[2],
		m[3], m[4], m[5],
		m[6], m[7], m[8])
}
