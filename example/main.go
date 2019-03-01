package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/xalanq/go-tracing/geo"
	"github.com/xalanq/go-tracing/geo/sphere"
	"github.com/xalanq/go-tracing/pic"
	"github.com/xalanq/go-tracing/ray"
	"github.com/xalanq/go-tracing/vec"
	"github.com/xalanq/go-tracing/world"
)

func main() {
	sample := 200
	if len(os.Args) > 1 {
		sample, _ = strconv.Atoi(os.Args[1])
	}
	zero := vec.NewZero()
	c1, c2 := vec.New(.75, .25, .25), vec.New(.25, .25, .75)
	c3, c4 := vec.New(.75, .75, .75), vec.Mult(vec.New(1, 1, 1), .999)
	p := pic.New(1024, 768)
	cam := ray.New(vec.New(50, 52, 295.6), vec.Norm(vec.New(0, -0.042612, -1)))
	depth := 5
	core := runtime.NumCPU()
	thread := 3
	world.New(cam, sample, depth, core, thread, 1.0, 1.5, 0.5135).
		Add(sphere.New(1e5, geo.New(vec.New(1e5+1, 40.8, 81.6), zero, c1, geo.Diffuse))).
		Add(sphere.New(1e5, geo.New(vec.New(-1e5+99, 40.8, 81.6), zero, c2, geo.Diffuse))).
		Add(sphere.New(1e5, geo.New(vec.New(50, 40.8, 1e5), zero, c3, geo.Diffuse))).
		Add(sphere.New(1e5, geo.New(vec.New(50, 40.8, -1e5+170), zero, zero, geo.Diffuse))).
		Add(sphere.New(1e5, geo.New(vec.New(50, 1e5, 81.6), zero, c3, geo.Diffuse))).
		Add(sphere.New(1e5, geo.New(vec.New(50, -1e5+81.6, 81.6), zero, c3, geo.Diffuse))).
		Add(sphere.New(16.5, geo.New(vec.New(27, 16.5, 47), zero, c4, geo.Specular))).
		Add(sphere.New(16.5, geo.New(vec.New(73, 16.5, 78), zero, c4, geo.Refractive))).
		Add(sphere.New(600, geo.New(vec.New(50, 681.6-.27, 81.6), vec.New(12, 12, 12), zero, geo.Diffuse))).
		Render(p)
	p.SavePPM(fmt.Sprintf("example_%v.ppm", sample))
}
