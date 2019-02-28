# Go Tracing

[![Build Status](https://travis-ci.org/xalanq/go-tracing.svg?branch=master)](https://travis-ci.org/xalanq/go-tracing)
[![Go Report Card](https://goreportcard.com/badge/github.com/xalanq/go-tracing)](https://goreportcard.com/report/github.com/xalanq/go-tracing)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.5-green.svg)](https://github.com/golang)
[![license](https://img.shields.io/badge/license-MIT-%23373737.svg)](https://raw.githubusercontent.com/xalanq/go-tracing/master/LICENSE)

A Go implement of path tracing and ray tracing in computer graphics.

# Feature

* Fast: Render with goroutine
* Structural: More OOP
* Progress bar: Use [cheggaaa/pb.v2](https://github.com/cheggaaa/pb/tree/v2)

# Usage

```Go
package main

import (
	"runtime"

	"github.com/xalanq/go-tracing/geo"
	"github.com/xalanq/go-tracing/geo/sphere"
	"github.com/xalanq/go-tracing/pic"
	"github.com/xalanq/go-tracing/ray"
	"github.com/xalanq/go-tracing/vec"
	"github.com/xalanq/go-tracing/world"
)

func main() {
	zero := vec.NewZero()
	c1, c2 := vec.New(.75, .25, .25), vec.New(.25, .25, .75)
	c3, c4 := vec.New(.75, .75, .75), vec.New(1, 1, 1).Mult(.999)
	p := pic.New(1024, 768)
	cam := ray.New(vec.New(50, 52, 295.6), vec.New(0, -0.042612, -1).Norm())
	sample := 200
	depth := 5
	thread := runtime.NumCPU() // no more than number of cpu core
	world.New(cam, sample, depth, thread, 1.0, 1.5, 0.5135).
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
	p.SavePPM("example_200.ppm")
}
```

# Reference

* [smallpt](http://www.kevinbeason.com/smallpt/)