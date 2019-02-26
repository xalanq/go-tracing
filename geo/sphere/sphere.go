package sphere

import (
	"math"

	"github.com/go-smallpt/geo"
	"github.com/go-smallpt/ray"
	"github.com/go-smallpt/vec"
)

// Sphere see smallpt
type Sphere struct {
	*geo.Geo
	R float64
}

// NewZero new zero sphere
func NewZero() *Sphere {
	return &Sphere{}
}

// New new one
func New(r float64, p, e, c *vec.Vec, t geo.ReflType) *Sphere {
	return &Sphere{
		R: r,
		Geo: &geo.Geo{
			Position: p,
			Emission: e,
			Color:    c,
			Type:     t,
		},
	}
}

// Hit intersect
func (s *Sphere) Hit(r *ray.Ray) float64 {
	op := s.Position.Sub(r.Origin)
	b := op.Dot(r.Direct)
	det := b*b - op.Len2() + s.R*s.R
	if det < 0.0 {
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
