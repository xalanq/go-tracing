package ray

import "github.com/xalanq/go-smallpt/vec"

// Ray ray
type Ray struct {
	Origin, Direct *vec.Vec
}

// New new one
func New(origin, direct *vec.Vec) *Ray {
	return &Ray{Origin: origin, Direct: direct}
}
