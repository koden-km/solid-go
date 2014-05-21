package solid

import (
	"encoding/json"
	"fmt"
	"math"
)

type Matrix4 [16]float64

func (m Matrix4) String() string {
	return fmt.Sprintf("%#v", m)
}

// TODO: take vector3's or vector4's?
//func NewMatrix4(a, b, c, d *Vector4) *Matrix4 {
//	return &Matrix4{
//		a.w, a.x, a.y, a.z,
//		b.w, b.x, b.y, b.z,
//		c.w, c.x, c.y, c.z,
//		a.w, a.x, a.y, a.z,
//	}
//}

func Vector4Identity() *Matrix4 {
	return &Matrix4{
		1, 0, 0, 0
		0, 1, 0, 0
		0, 0, 1, 0
		0, 0, 0, 1
	}
}

// TODO: add operation methods...
