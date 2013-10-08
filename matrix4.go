package solid

import (
	"fmt"
)

type Matrix4 [16]float32

func (m Matrix4) String() string {
	return fmt.Sprintf("m4{{%#v,%#v,%#v,%#v},{%#v,%#v,%#v,%#v},{%#v,%#v,%#v,%#v},{%#v,%#v,%#v,%#v}}",
		m[0], m[1], m[2], m[3],
		m[4], m[5], m[6], m[7],
		m[8], m[9], m[10], m[11],
		m[12], m[13], m[14], m[15])
}
