package world

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/xalanq/go-tracing/geo"
	"github.com/xalanq/go-tracing/pic"
	"github.com/xalanq/go-tracing/ray"
	"github.com/xalanq/go-tracing/vec"
	pb "gopkg.in/cheggaaa/pb.v1"
)

// World store Geo objects
type World struct {
	Objs             []geo.Hittable
	Cam              ray.Ray
	Sample, MaxDepth int
	Core, Thread     int
	Ratio            float64
	Na, Ng           float64 // refractive index of air and glass
	n1, n2           float64
}

// New new one (core is no more than runtime.NumCPU())
func New(cam ray.Ray, sample, maxDepth, core, thread int, na, ng, ratio float64) *World {
	if core > runtime.NumCPU() {
		core = runtime.NumCPU()
	}
	return &World{
		Cam:      cam,
		Sample:   sample,
		MaxDepth: maxDepth,
		Core:     core,
		Thread:   thread,
		Ratio:    ratio, Na: na, Ng: ng,
		n1: na / ng, n2: ng / na}
}

// Add add an object to the world
func (a *World) Add(obj geo.Hittable) *World {
	a.Objs = append(a.Objs, obj)
	return a
}

func (a *World) find(r ray.Ray) (obj geo.Hittable, g geo.Geo, pos, norm vec.Vec) {
	t := math.MaxFloat64
	for _, o := range a.Objs {
		if d := o.Hit(r); d != 0.0 && d < t {
			obj, t = o, d
		}
	}
	if obj != nil {
		g = obj.GetGeo()
		pos = vec.Add(r.Origin, vec.Mult(r.Direct, t))
		norm = vec.Norm(vec.Sub(pos, g.Position))
	}
	return
}

func max3(a, b, c float64) float64 {
	if a > b && a > c {
		return a
	}
	if b > c {
		return b
	}
	return c
}

func abs(x float64) float64 {
	if x < 0 {
		return x
	}
	return x
}

func (a *World) trace(r ray.Ray, depth int) vec.Vec {
	obj, g, pos, norm := a.find(r)
	if obj == nil {
		return vec.NewZero()
	}
	cl := g.Color.Copy()
	if depth++; depth > a.MaxDepth {
		if p := max3(cl.X, cl.Y, cl.Z); rand.Float64() < p {
			cl.Mult(1.0 / p)
		} else {
			return g.Emission.Copy()
		}
	}
	return vec.Add(g.Emission, vec.Mul(cl, func() vec.Vec {
		if g.Type == geo.Specular {
			d := vec.Sub(r.Direct, vec.Mult(norm, 2.0*norm.Dot(r.Direct)))
			return a.trace(ray.New(pos, d), depth)
		}
		w := norm
		if norm.Dot(r.Direct) >= 0.0 {
			w = vec.Mult(norm, -1)
		}
		if g.Type == geo.Diffuse {
			r1, r2 := 2.0*math.Pi*rand.Float64(), rand.Float64()
			r2s := math.Sqrt(r2)
			u := vec.New(1.0, 0.0, 0.0)
			if abs(w.X) > 0.1 {
				u.X, u.Y = 0.0, 1.0
			}
			u = vec.Norm(vec.Cross(u, w))
			v := vec.Cross(w, u)
			d := vec.Norm(vec.Add(vec.Add(vec.Mult(u, math.Cos(r1)*r2s), vec.Mult(v, math.Sin(r1)*r2s)), vec.Mult(w, math.Sqrt(1-r2))))
			return a.trace(ray.New(pos, d), depth)
		}
		refl := ray.New(pos, vec.Sub(r.Direct, vec.Mult(norm, 2.0*norm.Dot(r.Direct))))
		out, ddw, n, cos2t, sign := norm.Dot(w) <= 0.0, r.Direct.Dot(w), a.n1, 0.0, 1.0
		if out {
			n, sign = a.n2, -1.0
		}
		if cos2t = 1.0 - n*n*(1.0-ddw*ddw); cos2t < 0 {
			return a.trace(refl, depth)
		}
		td := vec.Norm(vec.Sub(vec.Mult(r.Direct, n), vec.Mult(norm, sign*(ddw*n+math.Sqrt(cos2t)))))
		refr := ray.New(pos, td)
		t1, t2 := a.Na-a.Ng, a.Na+a.Ng
		R0, c := (t1*t1)/(t2*t2), 1.0+ddw
		if out {
			c = 1.0 - td.Dot(norm)
		}
		Re := R0 + (1.0-R0)*c*c*c*c*c
		Tr := 1.0 - Re
		if depth > 2 {
			P := 0.25 + 0.5*Re
			if rand.Float64() < P {
				return vec.Mult(a.trace(refl, depth), Re/P)
			}
			return vec.Mult(a.trace(refr, depth), Tr/(1.0-P))
		}
		return vec.Add(vec.Mult(a.trace(refl, depth), Re), vec.Mult(a.trace(refr, depth), Tr))
	}()))
}

func gend() float64 {
	r := 2.0 * rand.Float64()
	if r < 1 {
		return math.Sqrt(r) - 1
	}
	return 1 - math.Sqrt(2-r)
}

type renderData struct {
	sum  *vec.Vec
	x, y float64
}

// Render render !!! store in p
func (a *World) Render(p *pic.Pic) *World {
	h, w := p.H, p.W
	fh, fw := float64(h), float64(w)
	cx := vec.New(fw*a.Ratio/fh, 0, 0)
	cy := vec.Mult(vec.Norm(vec.Cross(cx, a.Cam.Direct)), a.Ratio)
	sample := a.Sample / 4
	inv := 1.0 / float64(sample)

	fmt.Printf("w: %v, h: %v, sample: %v, actual sample: %v, thread: %v, cpu: %v\n", w, h, a.Sample, sample*4, a.Thread, a.Core)
	bar := pb.StartNew(h * w)
	bar.SetRefreshRate(1000 * time.Millisecond)

	runtime.GOMAXPROCS(a.Core)
	ch := make(chan renderData, a.Thread)
	wg := sync.WaitGroup{}
	wg.Add(a.Thread)

	for tid := 0; tid < a.Thread; tid++ {
		go func() {
			for {
				data, ok := <-ch
				if !ok {
					wg.Done()
					return
				}
				sum, x, y := data.sum, data.x, data.y
				for sy := 0.0; sy < 2.0; sy++ {
					for sx := 0.0; sx < 2.0; sx++ {
						c := vec.NewZero()
						for sp := 0; sp < sample; sp++ {
							ccx := vec.Mult(cx, ((sx+0.5+gend())/2.0+x)/fw-0.5)
							ccy := vec.Mult(cy, ((sy+0.5+gend())/2.0+y)/fh-0.5)
							d := vec.Add(vec.Add(ccx, ccy), a.Cam.Direct)
							r := ray.New(vec.Add(a.Cam.Origin, vec.Mult(d, 130)), vec.Norm(d))
							c.Add(vec.Mult(a.trace(r, 0), inv))
						}
						sum.Add(vec.Mult(vec.New(pic.Clamp(c.X), pic.Clamp(c.Y), pic.Clamp(c.Z)), 0.25))
					}
				}
				bar.Add(1)
			}
		}()
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			ch <- renderData{
				sum: &p.C[(h-y-1)*w+x],
				x:   float64(x),
				y:   float64(y),
			}
		}
	}

	close(ch)
	wg.Wait()

	bar.FinishPrint("Rendering completed")
	return a
}
