package sphere

import (
	"math"

	"github.com/go-smallpt/ray"
	"github.com/go-smallpt/vec"
)

// Refl see smallpt
type Refl int

const (
	// DIFF see smallpt
	DIFF Refl = iota
	// SPEC see smallpt
	SPEC
	// REFR see smallpt
	REFR
)

// Sphere see smallpt
type Sphere struct {
	rad     float64
	p, e, c *vec.Vec
	refl    Refl
}

// NewZero new zero sphere
func NewZero() *Sphere {
	return &Sphere{}
}

// New new one
func New(rad float64, p, e, c *vec.Vec, refl Refl) *Sphere {
	return &Sphere{rad: rad, p: p, e: e, c: c, refl: refl}
}

// Hit intersect
func (s *Sphere) Hit(r *ray.Ray) float64 {
	op := s.p.Sub(r.O)
	b := op.Dot(r.D)
	det := b*b - op.Len2() + s.rad*s.rad
	if det < 0 {
		return 0.0
	}
	det = math.Sqrt(det)
	eps := 1e-4
	if t := b - det; t > eps {
		return t
	}
	if t := b + det; t > eps {
		return t
	}
	return 0.0
}
