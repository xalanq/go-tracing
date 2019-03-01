package sphere

import (
	"math"

	"github.com/xalanq/go-tracing/geo"
	"github.com/xalanq/go-tracing/ray"
	"github.com/xalanq/go-tracing/vec"
)

// Sphere see smallpt
type Sphere struct {
	R float64
	geo.Geo
}

// New new one
func New(r float64, g geo.Geo) Sphere {
	return Sphere{
		R:   r,
		Geo: g,
	}
}

// Hit intersect
func (s Sphere) Hit(r ray.Ray) float64 {
	op := vec.Sub(s.Position, r.Origin)
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
