package geo

import (
	"github.com/xalanq/go-tracing/ray"
	"github.com/xalanq/go-tracing/vec"
)

// ReflType reflection type
type ReflType int

const (
	// Diffuse a type
	Diffuse ReflType = iota
	// Specular a type
	Specular
	// Refractive a type
	Refractive
)

// Geo Geometric object
type Geo struct {
	Position, Emission, Color *vec.Vec
	Type                      ReflType
}

// New new one
func New(position, emission, color *vec.Vec, tp ReflType) *Geo {
	return &Geo{
		Position: position,
		Emission: emission,
		Color:    color,
		Type:     tp,
	}
}

// GetGeo get geo itself
func (g *Geo) GetGeo() *Geo {
	return g
}

// Hittable hittable object
type Hittable interface {
	GetGeo() *Geo
	Hit(*ray.Ray) float64
}
