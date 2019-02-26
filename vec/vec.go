package vec

import (
	"fmt"
	"math"
)

// Vec 3d vector
type Vec struct {
	x, y, z float64
}

// NewZero new zero vec
func NewZero() *Vec {
	return &Vec{}
}

// New new vec
func New(x, y, z float64) *Vec {
	return &Vec{x: x, y: y, z: z}
}

// Copy copy one
func (a *Vec) Copy() *Vec {
	return &Vec{x: a.x, y: a.y, z: a.z}
}

func (a *Vec) String() string {
	return fmt.Sprintf("(%f, %f, %f)", a.x, a.y, a.z)
}

// Add a + b
func (a *Vec) Add(b *Vec) *Vec {
	a.x += b.x
	a.y += b.y
	a.z += b.z
	return a
}

// Add a + b (copy)
func Add(a, b *Vec) *Vec {
	return a.Copy().Add(b)
}

// Sub a - b
func (a *Vec) Sub(b *Vec) *Vec {
	a.x -= b.x
	a.y -= b.y
	a.z -= b.z
	return a
}

// Sub a - b (copy)
func Sub(a, b *Vec) *Vec {
	return a.Copy().Sub(b)
}

// Mul (a.x * b.x, a.y * b.y, a.z * b.z)
func (a *Vec) Mul(b *Vec) *Vec {
	a.x *= b.x
	a.y *= b.y
	a.z *= b.z
	return a
}

// Mul (a.x * b.x, a.y * b.y, a.z * b.z) (copy)
func Mul(a, b *Vec) *Vec {
	return a.Copy().Mul(b)
}

// Mult (a.x * t, a.y * t, a.z * t)
func (a *Vec) Mult(t float64) *Vec {
	a.x *= t
	a.y *= t
	a.z *= t
	return a
}

// Mult (a.x * t, a.y * t, a.z * t) (copy)
func Mult(a *Vec, t float64) *Vec {
	return a.Copy().Mult(t)
}

// Dot a.x*b.x + a.y*b.y + a.z*b.z
func Dot(a, b *Vec) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

// Dot a.x*b.x + a.y*b.y + a.z*b.z
func (a *Vec) Dot(b *Vec) float64 {
	return Dot(a, b)
}

// Len2 (length of a) ^ 2
func (a *Vec) Len2() float64 {
	return a.Dot(a)
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
func Norm(a *Vec) *Vec {
	return a.Copy().Norm()
}

// Cross cross (copy)
func Cross(a, b *Vec) *Vec {
	return &Vec{
		x: a.y*b.z - a.z*b.y,
		y: a.z*b.x - a.x*b.z,
		z: a.x*b.y - a.y*b.x,
	}
}

// Cross cross (copy)
func (a *Vec) Cross(b *Vec) *Vec {
	return Cross(a, b)
}

func clamp(x float64) float64 {
	if x < 0.0 {
		return 0.0
	} else if x > 1.0 {
		return 1.0
	}
	return x
}

// Clamp see smallpt
func (a *Vec) Clamp() *Vec {
	a.x = clamp(a.x)
	a.y = clamp(a.y)
	a.z = clamp(a.z)
	return a
}

// Clamp see smallpt (copy)
func Clamp(a *Vec) *Vec {
	return a.Copy().Clamp()
}

func toByte(x float64) byte {
	return byte(math.Pow(clamp(x), 1/2.2)*255 + 0.5)
}

// PPM ppm format string
func (a *Vec) PPM() string {
	return fmt.Sprintf("%v %v %v ", toByte(a.x), toByte(a.y), toByte(a.z))
}
