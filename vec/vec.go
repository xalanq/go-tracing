package vec

import (
	"fmt"
	"math"
)

// Vec 3d vector
type Vec struct {
	X, Y, Z float64
}

// NewZero new zero vec
func NewZero() Vec {
	return Vec{}
}

// New new vec
func New(x, y, z float64) Vec {
	return Vec{X: x, Y: y, Z: z}
}

// Copy copy one
func (a *Vec) Copy() Vec {
	return Vec{X: a.X, Y: a.Y, Z: a.Z}
}

func (a *Vec) String() string {
	return fmt.Sprintf("(%f, %f, %f)", a.X, a.Y, a.Z)
}

// Add a + b
func (a *Vec) Add(b Vec) *Vec {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
	return a
}

// Add a + b (copy)
func Add(a, b Vec) Vec {
	return *a.Add(b)
}

// Sub a - b
func (a *Vec) Sub(b Vec) *Vec {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z
	return a
}

// Sub a - b (copy)
func Sub(a, b Vec) Vec {
	return *a.Sub(b)
}

// Mul (a.X * b.X, a.Y * b.Y, a.Z * b.Z)
func (a *Vec) Mul(b Vec) *Vec {
	a.X *= b.X
	a.Y *= b.Y
	a.Z *= b.Z
	return a
}

// Mul (a.X * b.X, a.Y * b.Y, a.Z * b.Z) (copy)
func Mul(a, b Vec) Vec {
	return *a.Mul(b)
}

// Mult (a.X * t, a.Y * t, a.Z * t)
func (a *Vec) Mult(t float64) *Vec {
	a.X *= t
	a.Y *= t
	a.Z *= t
	return a
}

// Mult (a.X * t, a.Y * t, a.Z * t) (copy)
func Mult(a Vec, t float64) Vec {
	return *a.Mult(t)
}

// Dot a.X*b.X + a.Y*b.Y + a.Z*b.Z
func Dot(a, b Vec) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Dot a.X*b.X + a.Y*b.Y + a.Z*b.Z
func (a *Vec) Dot(b Vec) float64 {
	return Dot(*a, b)
}

// Len2 (length of a) ^ 2
func (a *Vec) Len2() float64 {
	return a.Dot(*a)
}

// Len length of a
func (a *Vec) Len() float64 {
	return math.Sqrt(a.Len2())
}

// Norm normalize length to 1
func (a *Vec) Norm() *Vec {
	return a.Mult(1.0 / a.Len())
}

// Norm normalize length to 1 (copy)
func Norm(a Vec) Vec {
	return *a.Norm()
}

// Cross cross (copy)
func Cross(a, b Vec) Vec {
	return Vec{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// Cross cross (copy)
func (a *Vec) Cross(b Vec) Vec {
	return Cross(*a, b)
}
