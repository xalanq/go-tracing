package geo

import (
	"github.com/go-smallpt/ray"
	"github.com/go-smallpt/vec"
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

// GetGeo get geo itself
func (g *Geo) GetGeo() *Geo {
	return g
}

// Hittable hittable object
type Hittable interface {
	GetGeo() *Geo
	Hit(*ray.Ray) float64
}
