package world

import (
	"github.com/xalanq/go-smallpt/geo"
	"github.com/xalanq/go-smallpt/pic"
)

// World store Geo objects
type World struct {
	objs []geo.Hittable
}

// New new one
func New() *World {
	return &World{}
}

// Add add an object to the world
func (a *World) Add(obj geo.Hittable) *World {
	a.objs = append(a.objs, obj)
	return a
}

// Render render !!! store in pic
func (a *World) Render(pic *pic.Pic) {

}
