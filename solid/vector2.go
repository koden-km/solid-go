package solid

import (
	"encoding/json"
	"fmt"
	"math"
)

// This type is based on github.com/cespare/frosty/vec.go
// Vector ops are modeled after math/big.

type Vector2 struct {
	X, Y float64
}

func (v *Vector2) String() string {
	return fmt.Sprintf("%#v", v)
}

func NewVector2(x, y float64) *Vector2 {
	return &Vector2{x, y}
}

func Vector2Zero() *Vector2 {
	return &Vector2{0, 0}
}

func Vector2One() *Vector2 {
	return &Vector2{1, 1}
}

func Vector2Up() *Vector2 {
	return &Vector2{0, 1}
}

func Vector2Forward() *Vector2 {
	return &Vector2{1, 0}
}

func (v *Vector2) Copy() *Vector2 {
	return &Vector2{v.X, v.Y}
}

func (v *Vector2) Add(a, b *Vector2) *Vector2 {
	v.X = a.X + b.X
	v.Y = a.Y + b.Y
	return v
}

func (v *Vector2) Sub(a, b *Vector2) *Vector2 {
	v.X = a.X - b.X
	v.Y = a.Y - b.Y
	return v
}

// Div sets v to be u / x for some scalar x and returns v.
func (v *Vector2) Div(u *Vector2, x float64) *Vector2 {
	v.X = u.X / x
	v.Y = u.Y / x
	return v
}

// Mul sets v to be u * x for some scalar x and returns v.
func (v *Vector2) Mul(u *Vector2, x float64) *Vector2 {
	v.X = u.X * x
	v.Y = u.Y * x
	return v
}

// Inv sets v to be the inverse of u and returns v.
func (v *Vector2) Inv(u *Vector2) *Vector2 {
	return v.Mul(v, -1)
}

// Mag returns the magnitude of v.
func (v *Vector2) Mag() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y))
}

// Normalize sets v to be the normalized (unit) vector of u and returns v.
func (v *Vector2) Normalize(u *Vector2) *Vector2 {
	return v.Div(u, u.Mag())
}

// Cross returns the cross product of u and v as a newly allocated vector.
// (This function does not follow the math/big pattern because it wouldn't work if the result vector were also
// one of the operands.)
func (v *Vector2) Cross(u *Vector2) *Vector2 {
	return &Vector2{
		X: (v.Y * u.Z) - (v.Z * u.Y),
		Y: (v.Z * u.X) - (v.X * u.Z),
	}
}

// Dot returns the dot product of u and v.
func (v *Vector2) Dot(u *Vector2) float64 {
	return (v.X * u.X) + (v.Y * u.Y)
}

func (v *Vector2) UnmarshalJSON(b []byte) error {
	a := [2]float64{}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	v.X, v.Y = a[0], a[1]
	return nil
}
