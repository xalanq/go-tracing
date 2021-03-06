# Go Tracing

[![Build Status](https://travis-ci.org/xalanq/go-tracing.svg?branch=master)](https://travis-ci.org/xalanq/go-tracing)
[![Go Report Card](https://goreportcard.com/badge/github.com/xalanq/go-tracing)](https://goreportcard.com/report/github.com/xalanq/go-tracing)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.6-green.svg)](https://github.com/golang)
[![license](https://img.shields.io/badge/license-MIT-%23373737.svg)](https://raw.githubusercontent.com/xalanq/go-tracing/master/LICENSE)

A Go implementation of path tracing in computer graphics.

# Feature

* Parallel: Render with goroutine
* Structural: More OOP
* Expandable: Easy to add a geometric object
* Progress bar: Use [cheggaaa/pb.v2](https://github.com/cheggaaa/pb/tree/v2)

# TO DO

* Wait...for a faster GO??????????
* Then....Add ray tracing and etc.

# Installation

`dep ensure -add github.com/xalanq/go-tracing`

or

`go get -u github.com/xalanq/go-tracing`

# Usage

```Go
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
```

# Reference

* [smallpt](http://www.kevinbeason.com/smallpt/)
